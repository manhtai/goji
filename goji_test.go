package goji

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestWithMatcher(t *testing.T) {
	exp := boolMatcher(true)
	ctx := WithMatcher(context.Background(), exp)
	if m := Match(ctx); m != exp {
		t.Errorf("expected %+v, got: %+v", exp, m)
	}
}

func TestWithMatcherRawPath(t *testing.T) {
	spec := NewPathSpec("/hello/:name")
	ctxSpec := WithMatcher(context.Background(), spec)
	m := Match(ctxSpec)

	if m != spec {
		t.Errorf("expected %+v, got: %+v", spec, m)
	}

	if fmt.Sprint(m) != "/hello/:name" {
		t.Errorf("expected /hello/:name, got: %+v", m)
	}
}

func TestWithHandler(t *testing.T) {
	exp := intHandler(1)
	ctx := WithHandler(context.Background(), exp)
	if h := ctx.Value(handlerKey).(http.Handler); h != exp {
		t.Errorf("expected %+v, got: %+v", exp, h)
	}
}

func TestWithPath(t *testing.T) {
	ctx := WithPath(context.Background(), "hi")
	if path := Path(ctx); path != "hi" {
		t.Errorf("expected hi, got: %q", path)
	}

	if path := Path(context.Background()); path != "" {
		t.Errorf("expected empty path, got: %q", path)
	}
}
