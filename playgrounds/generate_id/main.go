package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"

	gofrsUUID "github.com/gofrs/uuid/v5"
	googleUUID "github.com/google/uuid"
)

func main() {
	// https://github.com/google/uuid
	// https://pkg.go.dev/github.com/google/uuid
	fmt.Printf("github.com/google/uuid\n")
	fmt.Printf("\tuuid v1:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u1, _ := googleUUID.NewUUID()
			return u1.String()
		}())
	}
	fmt.Printf("\tuuid v2:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u2, _ := googleUUID.NewDCESecurity(googleUUID.Person, 1000)
			return u2.String()
		}())
	}
	fmt.Printf("\tuuid v3:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u3 := googleUUID.NewMD5(googleUUID.NameSpaceDNS, []byte("example.com"))
			return u3.String()
		}())
	}
	fmt.Printf("\tuuid v4:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u4, _ := googleUUID.NewRandom()
			return u4.String()
		}())
	}
	fmt.Printf("\tuuid v5:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u5 := googleUUID.NewSHA1(googleUUID.NameSpaceDNS, []byte("example.com"))
			return u5.String()
		}())
	}
	fmt.Printf("\tuuid v6:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u6, _ := googleUUID.NewV6()
			return u6.String()
		}())
	}
	fmt.Printf("\tuuid v7:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u7, _ := googleUUID.NewV7()
			return u7.String()
		}())
	}

	fmt.Println()

	// https://github.com/gofrs/uuid
	// https://pkg.go.dev/github.com/gofrs/uuid/v5
	fmt.Printf("github.com/gofrs/uuid\n")
	fmt.Printf("\tuuid v1:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u1, _ := gofrsUUID.NewV1()
			return u1.String()
		}())
	}
	fmt.Printf("\tuuid v2:\n")
	fmt.Printf("\t\t-\n")
	fmt.Printf("\tuuid v3:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u3 := gofrsUUID.NewV3(gofrsUUID.NamespaceDNS, "example.com")
			return u3.String()
		}())
	}
	fmt.Printf("\tuuid v4:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u4, _ := gofrsUUID.NewV4()
			return u4.String()
		}())
	}
	fmt.Printf("\tuuid v5:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u5 := gofrsUUID.NewV5(gofrsUUID.NamespaceDNS, "example.com")
			return u5.String()
		}())
	}
	fmt.Printf("\tuuid v6:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u6, _ := gofrsUUID.NewV6()
			return u6.String()
		}())
	}
	fmt.Printf("\tuuid v7:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			u7, _ := gofrsUUID.NewV7()
			return u7.String()
		}())
	}

	fmt.Println()

	// xid
	// https://github.com/rs/xid
	fmt.Printf("github.com/rs/xid\n")
	fmt.Printf("\txid:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			guid := xid.New()
			return guid.String()
		}())
	}

	fmt.Println()

	// ulid
	// https://github.com/oklog/ulid
	fmt.Printf("github.com/oklog/ulid\n")
	fmt.Printf("\tulid:\n")
	for range 10 {
		fmt.Printf("\t\t%s\n", func() string {
			return ulid.Make().String()
		}())
	}
}
