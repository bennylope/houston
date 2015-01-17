package main

import "testing"

func TestServices(t *testing.T) {

	var servicesList Services

	if len(servicesList.GetNames()) != 0 {
		t.Error("Empty Services list should be empty")
	}

	if len(servicesList.GetFiles()) != 0 {
		t.Error("Empty Services list should be empty")
	}

	servicesList.AddService(Service{"homebrew.mxcl.redis", "/Users/tester/Fake/homebrew.mxcl.redis.plist"})
	servicesList.AddService(Service{"com.spotify.webhelper", "/Users/tester/Fake/com.spotify.webhelper.plist"})
	servicesList.AddService(Service{"org.foo.bar", "/Users/tester/Fake/org.foo.bar.plist"})
	servicesList.AddService(Service{"org.virtualbox.vboxwebsrv", "/Users/tester/Fake/org.virtualbox.vboxwebsrv.plist"})

	if len(servicesList.GetFiles()) != 4 {
		t.Error("Service list should have 4 items")
	}

	filtered := servicesList.Filter("org")
	if len(filtered.GetNames()) != 2 {
		t.Error("Should be 2 filtered services")
	}

	redisList := servicesList.Filter("redis")
	if len(redisList.GetNames()) != 1 {
		t.Error("Should be 1 filtered service")
	}

	redis, err := servicesList.Get("redis")
	if err != nil {
		t.Error("redis service not found")
	}
	if redis.Name != "homebrew.mxcl.redis" {
		t.Error("redis service should have been returned, not ", redis.Name)
	}
}
