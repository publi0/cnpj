# cnpj

Consulte CNPJ pelo terminal

## Instalação

### From source

```bash
go build -o cnpj main.go
```

### From release

```bash
wget 
```

## Uso

```bash
./cnpj 00000000000000
```

## Saída

### Por padrão o output é o yaml simplificado

```yaml
cnpj_raiz: "19609281"
razao_social: ONEBOX TECNOLOGIA E INFORMACAO LTDA.
capital_social: "200000.00"
natureza_juridica:
    descricao: Sociedade Empresária Limitada
estabelecimento:
    cnpj: "19609281000102"
    nome_fantasia: ONEBOX TECNOLOGIA E INFORMACAO LTDA.
    situacao_cadastral: Ativa
    atividade_principal:
        descricao: Desenvolvimento de programas de computador sob encomenda
socios:
- nome: ALBERTO DE SOUSA LIMA
- nome: AUGUSTO CANIZELLA CHAGAS
- nome: BLUEBOX PARTICIPACOES E INVESTIMENTOS LTDA.
- nome: XGOLD PARTICIPACOES E INVESTIMENTOS LTDA.
- nome: VICTOR SHIMURA GOLDSZMIT
```

## Flags

### --ouput -o

Retorna o output em json ou yaml

```bash
./cnpj 00000000000000 --output json
```

### --raw -r

Retorna o output em texto puro

```bash
./cnpj 00000000000000 --raw
```

### --full -f

Retorna o output completo

```bash
./cnpj 00000000000000 --full
```

```yaml
cnpj_raiz: "19609281"
razao_social: ONEBOX TECNOLOGIA E INFORMACAO LTDA.
capital_social: "200000.00"
responsavel_federativo: ""
atualizado_em: 2024-06-08T03:00:00.000Z
porte:
  id: "05"
  descricao: Demais
natureza_juridica:
  id: "2062"
  descricao: Sociedade Empresária Limitada
qualificacao_do_responsavel:
  id: 5
  descricao: "Administrador "
socios:
- cpf_cnpj_socio: "***678648**"
  nome: ALBERTO DE SOUSA LIMA
  tipo: Pessoa Física
  data_entrada: 2014-01-27
  cpf_representante_legal: "***000000**"
  nome_representante: null
  faixa_etaria: 41 a 50 anos
  atualizado_em: 2024-06-08T03:00:00.000Z
  pais_id: "1058"
  qualificacao_socio:
    id: 5
    descricao: "Administrador "
  qualificacao_representante: null
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
- cpf_cnpj_socio: "***293098**"
  nome: AUGUSTO CANIZELLA CHAGAS
  tipo: Pessoa Física
  data_entrada: 2023-06-20
  cpf_representante_legal: "***000000**"
  nome_representante: null
  faixa_etaria: 41 a 50 anos
  atualizado_em: 2024-06-08T03:00:00.000Z
  pais_id: "1058"
  qualificacao_socio:
    id: 5
    descricao: "Administrador "
  qualificacao_representante: null
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
- cpf_cnpj_socio: "13833493000192"
  nome: BLUEBOX PARTICIPACOES E INVESTIMENTOS LTDA.
  tipo: Pessoa Jurídica
  data_entrada: 2014-01-27
  cpf_representante_legal: "***678648**"
  nome_representante: ALBERTO DE SOUSA LIMA
  faixa_etaria: Não se aplica
  atualizado_em: 2024-06-08T03:00:00.000Z
  pais_id: "1058"
  qualificacao_socio:
    id: 22
    descricao: Sócio
  qualificacao_representante:
    descricao: "Administrador "
    id: 5.0
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
- cpf_cnpj_socio: "27633935000161"
  nome: XGOLD PARTICIPACOES E INVESTIMENTOS LTDA.
  tipo: Pessoa Jurídica
  data_entrada: 2017-06-09
  cpf_representante_legal: "***293098**"
  nome_representante: AUGUSTO CANIZELLA CHAGAS
  faixa_etaria: Não se aplica
  atualizado_em: 2024-06-08T03:00:00.000Z
  pais_id: "1058"
  qualificacao_socio:
    id: 22
    descricao: Sócio
  qualificacao_representante:
    descricao: "Administrador "
    id: 5.0
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
- cpf_cnpj_socio: "***606638**"
  nome: VICTOR SHIMURA GOLDSZMIT
  tipo: Pessoa Física
  data_entrada: 2016-08-02
  cpf_representante_legal: "***000000**"
  nome_representante: null
  faixa_etaria: 41 a 50 anos
  atualizado_em: 2024-06-08T03:00:00.000Z
  pais_id: "1058"
  qualificacao_socio:
    id: 49
    descricao: "Sócio-Administrador "
  qualificacao_representante: null
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
simples: null
estabelecimento:
  cnpj: "19609281000102"
  atividades_secundarias:
  - id: "5320202"
    secao: H
    divisao: "53"
    grupo: "53.2"
    classe: 53.20-2
    subclasse: 5320-2/02
    descricao: Serviços de entrega rápida
  - id: "6202300"
    secao: J
    divisao: "62"
    grupo: "62.0"
    classe: 62.02-3
    subclasse: 6202-3/00
    descricao: Desenvolvimento e licenciamento de programas de computador customizáveis
  - id: "6204000"
    secao: J
    divisao: "62"
    grupo: "62.0"
    classe: 62.04-0
    subclasse: 6204-0/00
    descricao: Consultoria em tecnologia da informação
  - id: "6209100"
    secao: J
    divisao: "62"
    grupo: "62.0"
    classe: 62.09-1
    subclasse: 6209-1/00
    descricao: Suporte técnico, manutenção e outros serviços em tecnologia da informação
  - id: "6311900"
    secao: J
    divisao: "63"
    grupo: "63.1"
    classe: 63.11-9
    subclasse: 6311-9/00
    descricao: Tratamento de dados, provedores de serviços de aplicação e serviços de hospedagem na Internet
  - id: "6463800"
    secao: K
    divisao: "64"
    grupo: "64.6"
    classe: 64.63-8
    subclasse: 6463-8/00
    descricao: Outras sociedades de participação, exceto holdings
  - id: "7020400"
    secao: M
    divisao: "70"
    grupo: "70.2"
    classe: 70.20-4
    subclasse: 7020-4/00
    descricao: Atividades de consultoria em gestão empresarial, exceto consultoria técnica específica
  - id: "7490104"
    secao: M
    divisao: "74"
    grupo: "74.9"
    classe: 74.90-1
    subclasse: 7490-1/04
    descricao: Atividades de intermediação e agenciamento de serviços e negócios em geral, exceto imobiliários
  - id: "8299702"
    secao: "N"
    divisao: "82"
    grupo: "82.9"
    classe: 82.99-7
    subclasse: 8299-7/02
    descricao: Emissão de vales-alimentação, vales-transporte e similares
  - id: "8299799"
    secao: "N"
    divisao: "82"
    grupo: "82.9"
    classe: 82.99-7
    subclasse: 8299-7/99
    descricao: Outras atividades de serviços prestados principalmente às empresas não especificadas anteriormente
  cnpj_raiz: "19609281"
  cnpj_ordem: "0001"
  cnpj_digito_verificador: "02"
  tipo: Matriz
  nome_fantasia: ONEBOX TECNOLOGIA E INFORMACAO LTDA.
  situacao_cadastral: Ativa
  data_situacao_cadastral: 2014-01-27
  data_inicio_atividade: 2014-01-27
  nome_cidade_exterior: null
  tipo_logradouro: RUA
  logradouro: TENENTE NEGRAO
  numero: "140"
  complemento: CONJ  131                 SALA  01
  bairro: ITAIM BIBI
  cep: "04530030"
  ddd1: "11"
  telefone1: "43807757"
  ddd2: null
  telefone2: null
  ddd_fax: "11"
  fax: "32543333"
  email: ricardo@gimenessp.com.br
  situacao_especial: null
  data_situacao_especial: null
  atualizado_em: 2024-06-08T03:00:00.000Z
  atividade_principal:
    id: "6201501"
    secao: J
    divisao: "62"
    grupo: "62.0"
    classe: 62.01-5
    subclasse: 6201-5/01
    descricao: Desenvolvimento de programas de computador sob encomenda
  pais:
    id: "1058"
    iso2: BR
    iso3: BRA
    nome: Brasil
    comex_id: "105"
  estado:
    id: 26
    nome: São Paulo
    sigla: SP
    ibge_id: 35
  cidade:
    id: 3832
    nome: São Paulo
    ibge_id: 3550308
    siafi_id: "7107"
  motivo_situacao_cadastral: null
  inscricoes_estaduais: []
```

## Fonte de dados

URL base: <https://publica.cnpj.ws/cnpj/>
