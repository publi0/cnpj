package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type CNPJData struct {
	CnpjRaiz                  string           `json:"cnpj_raiz"`
	RazaoSocial               string           `json:"razao_social"`
	CapitalSocial             string           `json:"capital_social"`
	ResponsavelFederativo     string           `json:"responsavel_federativo"`
	AtualizadoEm              string           `json:"atualizado_em"`
	Porte                     Porte            `json:"porte"`
	NaturezaJuridica          NaturezaJuridica `json:"natureza_juridica"`
	QualificacaoDoResponsavel Qualificacao     `json:"qualificacao_do_responsavel"`
	Socios                    []Socio          `json:"socios"`
	Simples                   interface{}      `json:"simples"`
	Estabelecimento           Estabelecimento  `json:"estabelecimento"`
}

type Porte struct {
	Id        string `json:"id"`
	Descricao string `json:"descricao"`
}

type NaturezaJuridica struct {
	Id        string `json:"id"`
	Descricao string `json:"descricao"`
}

type Qualificacao struct {
	Id        int    `json:"id"`
	Descricao string `json:"descricao"`
}

type Socio struct {
	CpfCnpjSocio              string       `json:"cpf_cnpj_socio"`
	Nome                      string       `json:"nome"`
	Tipo                      string       `json:"tipo"`
	DataEntrada               string       `json:"data_entrada"`
	CpfRepresentanteLegal     string       `json:"cpf_representante_legal"`
	NomeRepresentante         interface{}  `json:"nome_representante"`
	FaixaEtaria               string       `json:"faixa_etaria"`
	AtualizadoEm              string       `json:"atualizado_em"`
	PaisId                    string       `json:"pais_id"`
	QualificacaoSocio         Qualificacao `json:"qualificacao_socio"`
	QualificacaoRepresentante interface{}  `json:"qualificacao_representante"`
	Pais                      Pais         `json:"pais"`
}

type Pais struct {
	Id      string `json:"id"`
	Iso2    string `json:"iso2"`
	Iso3    string `json:"iso3"`
	Nome    string `json:"nome"`
	ComexId string `json:"comex_id"`
}

type Estabelecimento struct {
	Cnpj                    string                `json:"cnpj"`
	AtividadesSecundarias   []AtividadeSecundaria `json:"atividades_secundarias"`
	CnpjRaiz                string                `json:"cnpj_raiz"`
	CnpjOrdem               string                `json:"cnpj_ordem"`
	CnpjDigitoVerificador   string                `json:"cnpj_digito_verificador"`
	Tipo                    string                `json:"tipo"`
	NomeFantasia            string                `json:"nome_fantasia"`
	SituacaoCadastral       string                `json:"situacao_cadastral"`
	DataSituacaoCadastral   string                `json:"data_situacao_cadastral"`
	DataInicioAtividade     string                `json:"data_inicio_atividade"`
	NomeCidadeExterior      interface{}           `json:"nome_cidade_exterior"`
	TipoLogradouro          string                `json:"tipo_logradouro"`
	Logradouro              string                `json:"logradouro"`
	Numero                  string                `json:"numero"`
	Complemento             string                `json:"complemento"`
	Bairro                  string                `json:"bairro"`
	Cep                     string                `json:"cep"`
	Ddd1                    string                `json:"ddd1"`
	Telefone1               string                `json:"telefone1"`
	Ddd2                    interface{}           `json:"ddd2"`
	Telefone2               interface{}           `json:"telefone2"`
	DddFax                  string                `json:"ddd_fax"`
	Fax                     string                `json:"fax"`
	Email                   string                `json:"email"`
	SituacaoEspecial        interface{}           `json:"situacao_especial"`
	DataSituacaoEspecial    interface{}           `json:"data_situacao_especial"`
	AtualizadoEm            string                `json:"atualizado_em"`
	AtividadePrincipal      AtividadeSecundaria   `json:"atividade_principal"`
	Pais                    Pais                  `json:"pais"`
	Estado                  Estado                `json:"estado"`
	Cidade                  Cidade                `json:"cidade"`
	MotivoSituacaoCadastral interface{}           `json:"motivo_situacao_cadastral"`
	InscricoesEstaduais     []interface{}         `json:"inscricoes_estaduais"`
}

type AtividadeSecundaria struct {
	Id        string `json:"id"`
	Secao     string `json:"secao"`
	Divisao   string `json:"divisao"`
	Grupo     string `json:"grupo"`
	Classe    string `json:"classe"`
	Subclasse string `json:"subclasse"`
	Descricao string `json:"descricao"`
}

type Estado struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Sigla  string `json:"sigla"`
	IbgeId int    `json:"ibge_id"`
}

type Cidade struct {
	Id      int    `json:"id"`
	Nome    string `json:"nome"`
	IbgeId  int    `json:"ibge_id"`
	SiafiId string `json:"siafi_id"`
}

const (
	CNPJURL = "https://publica.cnpj.ws/cnpj/"
)

func fetchCNPJData(ctx context.Context, cnpj string) (CNPJData, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CNPJURL+cnpj, nil)
	if err != nil {
		return CNPJData{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return CNPJData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CNPJData{}, err
	}

	var cnpjData CNPJData

	json.Unmarshal(body, &cnpjData)

	return cnpjData, nil
}
