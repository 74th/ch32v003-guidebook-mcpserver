package main

import (
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
)

type Loader struct {
	Manifest ManifestRoot
}

func loadDocumentManifest() ManifestRoot {
	var manifest ManifestRoot
	err := yaml.Unmarshal(DocumentManifestBlob, &manifest)
	if err != nil {
		panic(fmt.Sprintf("failed to load document manifest: %v", err))
	}

	return manifest
}

func NewLoader() *Loader {
	return &Loader{
		Manifest: loadDocumentManifest(),
	}
}

func (l *Loader) SelectManifestItem(name string) (*ManifestDocument, error) {
	for _, doc := range l.Manifest.Documents {
		if doc.Name == name {
			return &doc, nil
		}
	}
	return nil, fmt.Errorf("document %s not found", name)
}

func (l *Loader) GetMainMD(framework string) (string, error) {
	f, err := DocumentFS.Open("documents/" + framework + "/main.md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	out, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (l *Loader) GetItemMD(framework string, name string) (string, error) {
	m, err := l.SelectManifestItem(name)
	if err != nil {
		return "", err
	}

	f, err := DocumentFS.Open("documents/" + framework + "/" + m.MDFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	out, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (l *Loader) GetCodeDocument(framework string, name string) ([]string, error) {
	m, err := l.SelectManifestItem(name)
	if err != nil {
		return nil, err
	}

	var codes []string
	for _, cm := range m.Code[framework] {
		f, err := CodeFS.Open(cm.File)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		out, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		codes = append(codes, string(out))
	}

	return codes, nil
}
