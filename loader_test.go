package main

import "testing"

func TestLoader(t *testing.T) {
	loader := NewLoader()

	md, err := loader.GetMainMD("ch32fun")
	if err != nil {
		t.Fatalf("failed to get MD document: %v", err)
	}
	if md == "" {
		t.Fatal("MD document is empty")
	}

	md, err = loader.GetItemMD("ch32fun", "ADC")
	if err != nil {
		t.Fatalf("failed to get MD document: %v", err)
	}
	if md == "" {
		t.Fatal("MD document is empty")
	}

	codes, err := loader.GetCodeDocument("ch32fun", "ADC")
	if err != nil {
		t.Fatalf("failed to get code document: %v", err)
	}
	if len(codes) == 0 {
		t.Fatal("Code document is empty")
	}
	for _, code := range codes {
		if code == "" {
			t.Fatal("Code document is empty")
		}
	}
}

func TestAllMainMD(t *testing.T) {
	loader := NewLoader()

	frameworks := []string{"ch32fun", "wchsdk"}
	for _, framework := range frameworks {
		md, err := loader.GetMainMD(framework)
		if err != nil {
			t.Fatalf("failed to get MD document for %s: %v", framework, err)
		}
		if md == "" {
			t.Fatalf("MD document for %s is empty", framework)
		}
	}
}

func TestAllItemMD(t *testing.T) {
	loader := NewLoader()
	m := loader.Manifest

	frameworks := []string{"ch32fun", "wchsdk"}
	for _, framework := range frameworks {
		for _, doc := range m.Documents {
			md, err := loader.GetItemMD(framework, doc.Name)
			if err != nil {
				t.Fatalf("failed to get MD document for %s: %v", doc.Name, err)
			}
			if md == "" {
				t.Fatalf("MD document for %s is empty", doc.Name)
			}
		}
	}
}

func TestAllCode(t *testing.T) {
	loader := NewLoader()
	m := loader.Manifest

	frameworks := []string{"ch32fun", "wchsdk"}
	for _, framework := range frameworks {
		for _, doc := range m.Documents {
			codes, err := loader.GetCodeDocument(framework, doc.Name)
			if err != nil {
				t.Fatalf("failed to get code document for %s: %v", doc.Name, err)
			}
			if len(codes) != len(doc.Code[framework]) {
				t.Fatalf("Code document for %s is empty", doc.Name)
			}
			for _, code := range codes {
				if code == "" {
					t.Fatal("Code document is empty")
				}
			}
		}
	}
}
