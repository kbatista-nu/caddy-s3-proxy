package caddys3proxy

import (
	"fmt"
	"strconv"
	"time"

	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("s3proxy", parseCaddyfile)
}

// parseCaddyfile parses the s3proxy directive. It enables the proxying
// requests to S3 and configures it with this syntax:
//
//    s3proxy {
//	      region <aws region>
//	      bucket <s3 bucket name>
//	      index  <files...>
//	      TODO: browse [<template_file>]
//    }
//
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var b S3Proxy
	fmt.Printf("In Unmarshal")

	h.NextArg() // skip block beginning: "s3proxy"
	for h.NextBlock(0) {
		var err error
		switch h.Val() {
		//case "site_name":
		//err = parseStringArg(d, &b.SiteName)
		case "endpoint":
			err = parseStringArg(&h, &b.Endpoint)
		case "region":
			h.Args(&b.Region)
			// if b.Region == "" {
			// return nil, h.Err("region can not be empty")
			//}
			//case "key":
			//err = parseStringArg(d, &b.Key)
			//case "secret":
			//err = parseStringArg(d, &b.Secret)
		case "bucket":
			h.Args(&b.Bucket)
			if b.Bucket == "" {
				return nil, h.Err("bucket can not be empty")
			}
			//case "secure":
			//err = parseBoolArg(d, &b.Secure)
			//case "refresh_interval":
			//err = parseDurationArg(d, &b.RefreshInterval)
			//case "refresh_api_secret":
			//err = parseStringArg(d, &b.RefreshAPISecret)
			//case "debug":
			//err = parseBoolArg(d, &b.Debug)
			//case "signed_url_redirect":
			//err = parseBoolArg(d, &b.SignedURLRedirect)
		case "index":
			b.IndexNames = h.RemainingArgs()
			if len(b.IndexNames) == 0 {
				return nil, h.ArgErr()
			}
		default:
			err = h.Errf("%s not a valid s3proxy option", h.Val())
		}
		if err != nil {
			return nil, h.Errf("Error parsing %s: %s", h.Val(), err)
		}
	}

	return b, nil
}

func parseBoolArg(d *httpcaddyfile.Helper, out *bool) error {
	var strVal string
	err := parseStringArg(d, &strVal)
	if err == nil {
		*out, err = strconv.ParseBool(strVal)
	}
	return err
}

func parseDurationArg(d *httpcaddyfile.Helper, out *time.Duration) error {
	var strVal string
	err := parseStringArg(d, &strVal)
	if err == nil {
		*out, err = time.ParseDuration(strVal)
	}
	return err
}

func parseStringArg(d *httpcaddyfile.Helper, out *string) error {
	if !d.Args(out) {
		return d.ArgErr()
	}
	return nil
}