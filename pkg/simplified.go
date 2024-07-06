package pkg

type SimplifiedCNPJData struct {
	CnpjRaiz         string `json:"cnpj_raiz"`
	RazaoSocial      string `json:"razao_social"`
	CapitalSocial    string `json:"capital_social"`
	NaturezaJuridica struct {
		Descricao string `json:"descricao"`
	} `json:"natureza_juridica"`
	Estabelecimento struct {
		Cnpj               string `json:"cnpj"`
		NomeFantasia       string `json:"nome_fantasia"`
		SituacaoCadastral  string `json:"situacao_cadastral"`
		AtividadePrincipal struct {
			Descricao string `json:"descricao"`
		} `json:"atividade_principal"`
	} `json:"estabelecimento"`
	Socios []struct {
		Nome string `json:"nome"`
	} `json:"socios"`
}

func SimplifyCNPJData(data CNPJData) SimplifiedCNPJData {
	simplified := SimplifiedCNPJData{
		CnpjRaiz:      data.CnpjRaiz,
		RazaoSocial:   data.RazaoSocial,
		CapitalSocial: data.CapitalSocial,
	}
	simplified.NaturezaJuridica.Descricao = data.NaturezaJuridica.Descricao
	simplified.Estabelecimento.Cnpj = data.Estabelecimento.Cnpj
	simplified.Estabelecimento.NomeFantasia = data.Estabelecimento.NomeFantasia
	simplified.Estabelecimento.SituacaoCadastral = data.Estabelecimento.SituacaoCadastral
	simplified.Estabelecimento.AtividadePrincipal.Descricao = data.Estabelecimento.AtividadePrincipal.Descricao
	for _, socio := range data.Socios {
		simplified.Socios = append(simplified.Socios, struct {
			Nome string `json:"nome"`
		}{Nome: socio.Nome})
	}
	return simplified
}
