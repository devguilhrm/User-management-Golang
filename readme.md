# User Management CLI 🚀

Um aplicativo de linha de comando (CLI) em Go para gerenciamento de usuários, com persistência de dados em SQLite. Ideal para estudos e demonstração de arquitetura em camadas.

---

## 📋 Funcionalidades

- ✅ **Criar Usuário**: Cadastro de novos usuários com nome, e-mail e idade
- ✅ **Listar Usuários**: Visualiza todos os usuários cadastrados com contador
- ✅ **Atualizar Usuário**: Edita nome, e-mail ou idade de um usuário existente
- ✅ **Deletar Usuário**: Remove um usuário do banco de dados com confirmação
- ✅ **Persistência**: Dados salvos em banco SQLite (`db/users.db`)

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Descrição |
|-----------|-----------|
| **Go 1.21+** | Linguagem de programação principal |
| **SQLite3** | Banco de dados embutido para persistência |
| **google/uuid** | Geração de IDs únicos para cada usuário |
| **database/sql** | Interface padrão do Go para operações com banco de dados |

---

## 📁 Estrutura do Projeto

```
user-management/
├── main.go                 # Entry point e menu interativo
├── models/
│   └── user.go            # Definição da struct User
├── db/
│   └── connection.go      # Conexão e criação da tabela SQLite
├── repository/
│   └── user_repository.go # Operações CRUD no banco de dados
├── services/
│   └── user_service.go    # Regras de negócio (factory NewUser)
├── utils/
│   └── input.go           # Funções auxiliares para leitura de input
└── db/
    └── users.db           # Arquivo do banco SQLite (gerado automaticamente)
```

---

## 🚀 Como Executar

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado (versão 1.21 ou superior)
- GCC instalado (necessário para compilar o driver `sqlite3`)

### Passos

1. **Clone o repositório** (ou copie os arquivos para uma pasta):
```bash
mkdir user-management && cd user-management
# Copie os arquivos do projeto para esta pasta
```

2. **Inicialize o módulo Go**:
```bash
go mod init user-management
```

3. **Instale as dependências**:
```bash
go get github.com/google/uuid
go get github.com/mattn/go-sqlite3
```

4. **Execute o projeto**:
```bash
go run main.go
```

> 💡 **Dica**: Na primeira execução, o banco de dados `db/users.db` e a tabela `users` serão criados automaticamente.

---

## 🎮 Uso

Ao iniciar o programa, você verá o menu:

```
------- Cadastro de Usuários -------
(1) Criar Usuário
(2) Ver Usuários Cadastrados
(3) Atualizar Usuário
(4) Deletar Usuário
(5) Sair
```

### Exemplo de Fluxo

```bash
Escolha a opção: 1
Digite o Nome: Maria Silva
Digite o Email: maria@email.com
Digite a idade: 28
Usuário criado e cadastrado com sucesso!

Escolha a opção: 2
Usuário: (Maria Silva) | Email: (maria@email.com) | ID: (abc-123...) | Idade: (28)
O total de: 1 Cadastrados
```

---

## ⚙️ Detalhes Técnicos

### Modelo de Dados (`User`)
```go
type User struct {
    ID    string  // UUID gerado automaticamente
    Name  string
    Email string
    Age   int
}
```

### Camadas da Aplicação

| Camada | Responsabilidade |
|--------|-----------------|
| `main` | Interface com usuário e orquestração |
| `services` | Regras de negócio e criação de entidades |
| `repository` | Acesso e operações no banco de dados |
| `models` | Definição de estruturas de dados |
| `db` | Configuração de conexão e schema |
| `utils` | Funções utilitárias de I/O |

### Tratamento de Erros
- Uso de `log.Fatal` para erros críticos de conexão/BD
- Mensagens amigáveis para erros de negócio (ex: usuário não encontrado)
- Fechamento adequado de recursos com `defer`

---

## 🧪 Melhorias Sugeridas

- [ ] Adicionar validações de e-mail e idade
- [ ] Implementar busca por ID ou e-mail (além do nome)
- [ ] Substituir `log.Fatal` por tratamento de erro mais granular
- [ ] Adicionar testes unitários para `repository` e `services`
- [ ] Implementar paginação na listagem de usuários
- [ ] Criar script de migração para versionamento do schema

---

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## 📄 Licença

Este projeto é de caráter educacional e está sob a licença MIT. Sinta-se à vontade para usar, modificar e distribuir.

---

> 💡 **Nota para Desenvolvedores**: Este projeto segue uma arquitetura simples em camadas, ideal para aprendizado de Go, padrões de repositório e interação com bancos de dados relacionais.

---

*Desenvolvido com ❤️ em Go* 🐹