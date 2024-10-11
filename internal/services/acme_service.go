/*
Package services provides ACME (Let's Encrypt) integration using the lego library.
*/

package services

import (
	"crypto"
	"crypto/tls"
	"io/ioutil"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/registration"
)

// ACMEClient manages Let's Encrypt certificates using the ACME protocol.
type ACMEClient struct {
	config *lego.Config
	client *lego.Client
}

// User implements the acme.User interface required by the lego library.
type User struct {
	Email        string
	Registration *registration.Resource
	Key          crypto.PrivateKey // Changed to use the crypto package
}

// GetEmail returns the user's email address.
func (u *User) GetEmail() string {
	return u.Email
}

// GetRegistration returns the user's registration resource.
func (u *User) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey returns the user's private key.
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

// NewACMEClient creates a new ACMEClient for the specified email and domain.
func NewACMEClient(email, domain string) (*ACMEClient, error) {
	privateKey, err := certcrypto.GeneratePrivateKey(certcrypto.RSA2048)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email: email,
		Key:   privateKey,
	}

	acmeCfg := lego.NewConfig(user)
	acmeCfg.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(acmeCfg)
	if err != nil {
		return nil, err
	}

	// Configure DNS provider (e.g., Cloudflare)
	provider, err := cloudflare.NewDNSProvider()
	if err != nil {
		return nil, err
	}

	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		return nil, err
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return nil, err
	}
	user.Registration = reg

	return &ACMEClient{
		config: acmeCfg,
		client: client,
	}, nil
}

// ObtainCertificate generates a TLS certificate for the specified domain using Let's Encrypt.
func (ac *ACMEClient) ObtainCertificate(domain string) (*tls.Certificate, error) {
	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}

	certificates, err := ac.client.Certificate.Obtain(request)
	if err != nil {
		return nil, err
	}

	// Save certificates to disk
	err = ioutil.WriteFile(domain+".crt", certificates.Certificate, 0600)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(domain+".key", certificates.PrivateKey, 0600)
	if err != nil {
		return nil, err
	}

	// Load the certificate
	tlsCert, err := tls.LoadX509KeyPair(domain+".crt", domain+".key")
	if err != nil {
		return nil, err
	}

	return &tlsCert, nil
}
