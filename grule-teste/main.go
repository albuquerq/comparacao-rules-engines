package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/hyperjumptech/grule-rule-engine/engine"

	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type PessoaFato struct {
	Nome string
}

func (pf PessoaFato) Ola() {
	fmt.Println("Olá mundo! Meu nome é " + pf.Nome + "!")
}

type lib struct {
}

func (lib) CumprimenteZe() {
	fmt.Println("Aeeeê! É um zé!!!!")
}

func (lib) IsZe(name string) bool {
	return strings.Contains(name, "Zé")
}

func main() {
	contexto := ast.NewDataContext()

	err := contexto.Add("pessoa", &PessoaFato{Nome: "Zé Ninguém"})
	if err != nil {
		log.Fatalln("Erro ao adicionar dados ao contexto", err)
	}

	err = contexto.Add("lib", &lib{})
	if err != nil {
		log.Fatal(err)
	}

	libConhecimento := ast.NewKnowledgeLibrary()

	builder := builder.NewRuleBuilder(libConhecimento)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	regras := pkg.NewFileResource(path.Join(cwd, "regras", "rules.grl"))

	err = builder.BuildRuleFromResource("BasePessoa", "0.0.1", regras)
	if err != nil {
		log.Fatal("Erro ao construir as regras a partir do arquivo", err)
	}

	eng := engine.NewGruleEngine()

	baseConhecimento := libConhecimento.NewKnowledgeBaseInstance("BasePessoa", "0.0.1")

	err = eng.Execute(contexto, baseConhecimento)
	if err != nil {
		log.Fatalln("Erro ao executar regras", err)
	}
}
