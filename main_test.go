package main

import (
	"io/ioutil"
	"path"
	"reflect"
	"testing"
)

func Test_ensureFastlaneVersionAndCreateCmdSlice(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("init: failed tp create temp dir")
	}

	gemfileLockPath := path.Join(tempDir, "Gemfile.lock")
	gemfilePath := path.Join(tempDir, "Gemfile")
	if err := ioutil.WriteFile(gemfileLockPath, []byte(gemfileLock), 0500); err != nil {
		t.Errorf("init: falied to create Gemfile.lock")
	}

	if err := ioutil.WriteFile(gemfilePath, []byte(gemfile), 0500); err != nil {
		t.Errorf("init: falied to create Gemfile")
	}

	tests := []struct {
		name         string
		forceVersion string
		gemfilePth   string
		want         []string
		want1        string
		wantErr      bool
	}{
		{
			name:       "test bundler install",
			gemfilePth: gemfilePath,
			want:       []string{"bundle", "_2.0.2_", "exec", "fastlane"},
			want1:      tempDir,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ensureFastlaneVersionAndCreateCmdSlice(tt.forceVersion, tt.gemfilePth)
			if (err != nil) != tt.wantErr {
				t.Errorf("ensureFastlaneVersionAndCreateCmdSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ensureFastlaneVersionAndCreateCmdSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ensureFastlaneVersionAndCreateCmdSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

const gemfile = `
source 'https://rubygems.org'

gem 'fastlane', '~> 2.123.0'
`

const gemfileLock = `
GEM
  remote: https://rubygems.org/
  specs:
    CFPropertyList (3.0.0)
    addressable (2.6.0)
      public_suffix (>= 2.0.2, < 4.0)
    atomos (0.1.3)
    babosa (1.0.2)
    claide (1.0.2)
    colored (1.2)
    colored2 (3.1.2)
    commander-fastlane (4.4.6)
      highline (~> 1.7.2)
    declarative (0.0.10)
    declarative-option (0.1.0)
    digest-crc (0.4.1)
    domain_name (0.5.20190701)
      unf (>= 0.0.5, < 1.0.0)
    dotenv (2.7.4)
    emoji_regex (1.0.1)
    excon (0.64.0)
    faraday (0.15.4)
      multipart-post (>= 1.2, < 3)
    faraday-cookie_jar (0.0.6)
      faraday (>= 0.7.4)
      http-cookie (~> 1.0.0)
    faraday_middleware (0.13.1)
      faraday (>= 0.7.4, < 1.0)
    fastimage (2.1.5)
    fastlane (2.123.0)
      CFPropertyList (>= 2.3, < 4.0.0)
      addressable (>= 2.3, < 3.0.0)
      babosa (>= 1.0.2, < 2.0.0)
      bundler (>= 1.12.0, < 3.0.0)
      colored
      commander-fastlane (>= 4.4.6, < 5.0.0)
      dotenv (>= 2.1.1, < 3.0.0)
      emoji_regex (>= 0.1, < 2.0)
      excon (>= 0.45.0, < 1.0.0)
      faraday (~> 0.9)
      faraday-cookie_jar (~> 0.0.6)
      faraday_middleware (~> 0.9)
      fastimage (>= 2.1.0, < 3.0.0)
      gh_inspector (>= 1.1.2, < 2.0.0)
      google-api-client (>= 0.21.2, < 0.24.0)
      google-cloud-storage (>= 1.15.0, < 2.0.0)
      highline (>= 1.7.2, < 2.0.0)
      json (< 3.0.0)
      mini_magick (~> 4.5.1)
      multi_json
      multi_xml (~> 0.5)
      multipart-post (~> 2.0.0)
      plist (>= 3.1.0, < 4.0.0)
      public_suffix (~> 2.0.0)
      rubyzip (>= 1.2.2, < 2.0.0)
      security (= 0.1.3)
      simctl (~> 1.6.3)
      slack-notifier (>= 2.0.0, < 3.0.0)
      terminal-notifier (>= 2.0.0, < 3.0.0)
      terminal-table (>= 1.4.5, < 2.0.0)
      tty-screen (>= 0.6.3, < 1.0.0)
      tty-spinner (>= 0.8.0, < 1.0.0)
      word_wrap (~> 1.0.0)
      xcodeproj (>= 1.8.1, < 2.0.0)
      xcpretty (~> 0.3.0)
      xcpretty-travis-formatter (>= 0.0.3)
    gh_inspector (1.1.3)
    google-api-client (0.23.9)
      addressable (~> 2.5, >= 2.5.1)
      googleauth (>= 0.5, < 0.7.0)
      httpclient (>= 2.8.1, < 3.0)
      mime-types (~> 3.0)
      representable (~> 3.0)
      retriable (>= 2.0, < 4.0)
      signet (~> 0.9)
    google-cloud-core (1.3.0)
      google-cloud-env (~> 1.0)
    google-cloud-env (1.2.0)
      faraday (~> 0.11)
    google-cloud-storage (1.16.0)
      digest-crc (~> 0.4)
      google-api-client (~> 0.23)
      google-cloud-core (~> 1.2)
      googleauth (>= 0.6.2, < 0.10.0)
    googleauth (0.6.7)
      faraday (~> 0.12)
      jwt (>= 1.4, < 3.0)
      memoist (~> 0.16)
      multi_json (~> 1.11)
      os (>= 0.9, < 2.0)
      signet (~> 0.7)
    highline (1.7.10)
    http-cookie (1.0.3)
      domain_name (~> 0.5)
    httpclient (2.8.3)
    json (2.2.0)
    jwt (2.2.1)
    memoist (0.16.0)
    mime-types (3.2.2)
      mime-types-data (~> 3.2015)
    mime-types-data (3.2019.0331)
    mini_magick (4.5.1)
    multi_json (1.13.1)
    multi_xml (0.6.0)
    multipart-post (2.0.0)
    nanaimo (0.2.6)
    naturally (2.2.0)
    os (1.0.1)
    plist (3.5.0)
    public_suffix (2.0.5)
    representable (3.0.4)
      declarative (< 0.1.0)
      declarative-option (< 0.2.0)
      uber (< 0.2.0)
    retriable (3.1.2)
    rouge (2.0.7)
    rubyzip (1.2.3)
    security (0.1.3)
    signet (0.11.0)
      addressable (~> 2.3)
      faraday (~> 0.9)
      jwt (>= 1.5, < 3.0)
      multi_json (~> 1.10)
    simctl (1.6.5)
      CFPropertyList
      naturally
    slack-notifier (2.3.2)
    terminal-notifier (2.0.0)
    terminal-table (1.8.0)
      unicode-display_width (~> 1.1, >= 1.1.1)
    tty-cursor (0.7.0)
    tty-screen (0.7.0)
    tty-spinner (0.9.1)
      tty-cursor (~> 0.7)
    uber (0.1.0)
    unf (0.1.4)
      unf_ext
    unf_ext (0.0.7.6)
    unicode-display_width (1.6.0)
    word_wrap (1.0.0)
    xcodeproj (1.10.0)
      CFPropertyList (>= 2.3.3, < 4.0)
      atomos (~> 0.1.3)
      claide (>= 1.0.2, < 2.0)
      colored2 (~> 3.1)
      nanaimo (~> 0.2.6)
    xcpretty (0.3.0)
      rouge (~> 2.0.7)
    xcpretty-travis-formatter (1.0.0)
      xcpretty (~> 0.2, >= 0.0.7)

PLATFORMS
  ruby

DEPENDENCIES
  fastlane (~> 2.123.0)

BUNDLED WITH
   2.0.2
`
