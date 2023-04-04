# UNION RESTful API

<div align="center">
  <img src="https://img.shields.io/static/v1?label=golang&message=language&color=blue&style=for-the-badge&logo=GO"/>
  <img src="https://img.shields.io/static/v1?label=gin%20gonic&message=framework&color=orange&style=for-the-badge&logo=GO"/>
  <img src="http://img.shields.io/static/v1?label=License&message=MIT&color=green&style=for-the-badge"/>
  <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/ViniciusGR797/crm_union?style=for-the-badge">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/ViniciusGR797/crm_union?style=for-the-badge">
  <img alt="GitHub forks" src="https://img.shields.io/github/forks/ViniciusGR797/crm_union?style=for-the-badge">
  <img alt="Bitbucket open issues" src="https://img.shields.io/bitbucket/issues/ViniciusGR797/crm_union?style=for-the-badge">
  <img alt="Bitbucket open pull request" src="https://img.shields.io/bitbucket/pr-raw/ViniciusGR797/crm_union?style=for-the-badge">
  <img src="http://img.shields.io/static/v1?label=STATUS&message=Development&color=GREEN&style=for-the-badge"/>
  
</div>

<div align="center">
  <img src="https://cdn.discordapp.com/attachments/1089358473483006105/1092501283291795527/LogoUnion.png" alt="logo HappyFit">
</div>
  
> UNION é um software de CRM que auxilia a empresa no gerenciamento de relacionamento com clientes, permitindo a organização e análise de dados de clientes para melhorar as estratégias de vendas e marketing.

### Tópicos 

:small_blue_diamond: [⚙ Execução localmente](#-execução-localmente)

:small_blue_diamond: [🐳 Execução com Docker](#-execução-com-docker)

:small_blue_diamond: [📃 Executando os testes unitários](#-executando-os-testes-unitários)

:small_blue_diamond: [🛠 Construído com](#-construído-com)

:small_blue_diamond: [👥 Equipe](#-equipe)

:small_blue_diamond: [📄 Licença](#-licença)

:small_blue_diamond: [🎁 Agradecimentos](#-agradecimentos)

### ⚙ Execução localmente

Para executar o projeto localmente na sua máquina, basta dar o seguinte comando no terminal na pasta raiz do projeto:

```
go run main.go
```

Para executar um micro serviço específico:

```
go run microservice_user/cmd/main.go
```

### 🐳 Execução com Docker
Para executar o projeto com docker, basta dar o seguinte comando no terminal na pasta raiz do projeto:
```
docker-compose -p crm-union up
```

## 📃 Executando os testes unitários

Entre na pasta do microserviço que deseja testar, por exemplo:
```
cd microservice_user
```

Execute os testes de todo microserviço:
```
go test ./... -v -cover
```

Para salvar o resultado dos testes em um JSON e ter mais detalhes sobre os teste:
```
go test ./... -v -cover -json > results.json
```

## 🛠 Construído com
Foi utilizado as seguintes tecnologias para construir o software:

* [Golang](https://go.dev/doc/) - linguagem de programação
* [Gin Gonic](https://github.com/gin-gonic/gin) - framework web 
* [MySQL](https://dev.mysql.com/doc/) - banco de dados

## 👥 Equipe

Segue abaixo a lista de pessoas a quem expressamos nossa gratidão pela valiosa contribuição para o sucesso deste projeto.

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Breno-AA">
        <img src="https://avatars.githubusercontent.com/u/45759261?v=4" width="100px;" alt="Foto do Breno"/><br>
        <sub>
          <b>Breno de Ávila Almeida</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Eduardo-Lisboa">
        <img src="https://avatars.githubusercontent.com/u/108145431?v=4" width="100px;" alt="Foto do Eduardo"/><br>
        <sub>
          <b>Eduardo Henrique Lisboa Alves</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/erich-ika">
        <img src="https://avatars.githubusercontent.com/u/95936922?v=4" width="100px;" alt="Foto do Iago"/><br>
        <sub>
          <b>Erich Iago Knoor de Almeida</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Guilhermealst">
        <img src="https://avatars.githubusercontent.com/u/116703599?v=4" width="100px;" alt="Foto do Guilherme"/><br>
        <sub>
          <b>Guilherme Alvarenga Morassutti</b>
        </sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/Jeffersonbortoluzzi">
        <img src="https://avatars.githubusercontent.com/u/95036870?v=4" width="100px;" alt="Foto do Jefferson"/><br>
        <sub>
          <b>Jefferson Bortoluzzi de Oliveira Gordo</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/JulioFavaro98">
        <img src="https://avatars.githubusercontent.com/u/94727206?v=4" width="100px;" alt="Foto do Julio"/><br>
        <sub>
          <b>Julio dos Santos Fávaro</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/leohlima">
        <img src="https://avatars.githubusercontent.com/u/88112020?v=4" width="100px;" alt="Foto do Leonardo"/><br>
        <sub>
          <b>Leonardo Lima de Araujo</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/LuisFelipeRodrigo">
        <img src="https://avatars.githubusercontent.com/u/103063554?v=4" width="100px;" alt="Foto do Luis Felipe"/><br>
        <sub>
          <b>Luis Felipe Rodrigo da Silva</b>
        </sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/GuilhermeFegueredo">
        <img src="https://avatars.githubusercontent.com/u/80172606?v=4" width="100px;" alt="Foto do Luis Guilherme"/><br>
        <sub>
          <b>Luis Guilherme Fegueredo de Oliveira</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/MarianySecundini">
        <img src="https://avatars.githubusercontent.com/u/86186072?v=4" width="100px;" alt="Foto da Mariany"/><br>
        <sub>
          <b>Mariany Secundini de Almeida Siqueira</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Raytsson">
        <img src="https://avatars.githubusercontent.com/u/107050617?v=4" width="100px;" alt="Foto do Raytsson"/><br>
        <sub>
          <b>Raytsson Martini de Moraes</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/eu-micaeu">
        <img src="https://avatars.githubusercontent.com/u/69124656?v=4" width="100px;" alt="Foto do Micael"/><br>
        <sub>
          <b>Micael Rocha</b>
        </sub>
      </a>
    </td>
  </tr>
  <tr>
    <td align="center">
      <a href="https://github.com/taglyscostacurta">
        <img src="https://avatars.githubusercontent.com/u/97138443?v=4" width="100px;" alt="Foto do Taglys"/><br>
        <sub>
          <b>Taglys Henrique Costacurta de Oliveira</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Avelar19">
        <img src="https://avatars.githubusercontent.com/u/118865712?v=4" width="100px;" alt="Foto do Thiago"/><br>
        <sub>
          <b>Thiago Augusto Avelar da Silva</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/V1n1c1usAu">
        <img src="https://avatars.githubusercontent.com/u/96804098?v=4" width="100px;" alt="Foto do Luiz"/><br>
        <sub>
          <b>Vinícius Augusto Dos Santos Silva</b>
        </sub>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/ViniciusGR797">
        <img src="https://avatars.githubusercontent.com/u/106624536?v=4" width="100px;" alt="Foto do Vinícius"/><br>
        <sub>
          <b>Vinícius Gomes Ribeiro</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

## 📄 Licença

Este projeto está sob a licença (MIT License) - veja o arquivo [LICENSE.md](https://github.com/ViniciusGR797/crm_union/blob/main/license) para detalhes.

## 🎁 Agradecimentos

* Gilberto Anderson (Giba)

[⬆ Voltar ao topo](#union-restful-api)<br>
