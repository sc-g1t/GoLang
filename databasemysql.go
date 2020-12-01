// Pacote do programa.
package main

// Bibliotecas
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type categoriaa struct {
	id   int
	nome string
}

type marcaa struct {
	id   int
	nome string
}

// ???
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Abre e loga como root no mysql
	db, err := sql.Open("mysql", "root:abcd1234@/")
	if err != nil {
		panic(err)
	}

	// Fecha o Banco de Dados após a execução
	defer db.Close()

	// Cria Database webloja se já não existir
	fmt.Println("\n\nCriando banco de dados webloja.")
	exec(db, "CREATE DATABASE IF NOT EXISTS webloja")

	// Escolhe o Database webloja para ser usado
	fmt.Println("Selecionando banco de dados webloja.")
	exec(db, "USE webloja")

	// Criando tabelas

	// Cria tabela Lojas

	fmt.Println("Criando Tabela Lojas.")
	exec(db, `CREATE TABLE IF NOT EXISTS lojas(
		loja_id INT AUTO_INCREMENT PRIMARY KEY,
		loja VARCHAR(40),
		cidadeloja VARCHAR(40),
		estadoloja VARCHAR(40)
	)`)

	// Cria tabela Produtos

	fmt.Println("Criando Tabela Produtos.")
	exec(db, `CREATE TABLE IF NOT EXISTS produtos(
		produto_id INT AUTO_INCREMENT PRIMARY KEY,
		produto VARCHAR(40),
		quantidade INT NOT NULL,
		vencimento DATE NOT NULL,
		preco DECIMAL(20,2),
		custo DECIMAL(20,2)
	)`)

	// Cria tabela Estoque
	fmt.Println("Criando Tabela Estoque.")
	exec(db, `CREATE TABLE IF NOT EXISTS estoques(
		loja_id INT NOT NULL,
		produto_id INT NOT NULL,
		quantidade INT NOT NULL,
		PRIMARY KEY(loja_id, produto_id),
		FOREIGN KEY (loja_id) REFERENCES lojas(loja_id),
		FOREIGN KEY (produto_id) REFERENCES produtos(produto_id)
	)
`)

	// Cria tabela Categorias.
	fmt.Println("Criando tabela categorias.")
	exec(db, `CREATE TABLE IF NOT EXISTS categorias(
		categoria_id INT AUTO_INCREMENT PRIMARY KEY,
		categoria_nome VARCHAR(40)
	)`)

	// Cria tabela Marcas.
	fmt.Println("Criando tabela marcas.")
	exec(db, `CREATE TABLE IF NOT EXISTS marcas(
		marca_id INT AUTO_INCREMENT PRIMARY KEY,
		marca_nome VARCHAR(40)
	)`)

	// Declaração de variáveis
	var categorias []string
	var marcas []string
	marca := ""
	categoria := ""
	resposta := ""
	nome := ""
	totalregistros := 0
	codigo := 0

	// Boas-vindas
	fmt.Println("\n\n------ Seja bem-vindo ! ------")
	// Menu
	for resposta != "0" {
		fmt.Println("\n\n	------ MENU ------")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("\nPara criar uma nova Categoria digite 1")
		fmt.Println("Para cirar uma nova Marca digite 2")
		fmt.Println("Para mostar o conteudo da tabela Categoria digite 3")
		fmt.Println("Para mostrar o conteudo da tabela Marcas digite 4")
		fmt.Println("Para atualizar uma Categoria digite 5")
		fmt.Println("Para atualizar uma Marca digite 6")
		fmt.Println("Para DELETAR a tabela Categorias digite 7")
		fmt.Println("Para DELETAR a tabela Marcas digite 8")
		fmt.Println("Para DELETAR as tabelas Marcas e Categorias digite 9")
		fmt.Println("Para SAIR digite 0")
		fmt.Scanln(&resposta)

		switch resposta {

		case "1":
			// Cria uma nova catégoria (Input do usuario)
			fmt.Println("\nDeseja criar uma nova categoria? s/n")
			fmt.Scanln(&resposta)

			if resposta == "s" {
				codigo := 0
				for {
					fmt.Println("Informe um nome para a categoria: ")
					fmt.Scanln(&categoria)
					categorias = append(categorias, categoria)

					fmt.Println("Deseja cadastrar outra categoria? s/n")
					fmt.Scanln(&resposta)

					if resposta != "s" {
						break
					}
				}

				// Seleciona o max(categoria_id) como tamanho do código + executa db.Query que acessa o BD
				rows, _ := db.Query("select max(categoria_id) as ultimo from categorias")
				defer rows.Close()
				for rows.Next() {
					rows.Scan(&codigo)
					codigo = codigo + 1
				}
				totalregistros = len(categorias)
				tx, _ := db.Begin()
				stmt, _ := tx.Prepare("insert into categorias(categoria_id, categoria_nome) values(?,?)")

				// ??? Faz o incremento no categoria_id ???
				for i := 0; i < totalregistros; i++ {
					stmt.Exec(codigo, categorias[i])
					codigo = codigo + 1

				}

				// ???
				if err != nil {
					tx.Rollback()
					log.Fatal(err)
				}
				tx.Commit()
			}

		case "2":
			// Cria uma nova marca (Input do usuario)

			fmt.Println("Gostaria de adicionar uma nova marca? s/n")
			fmt.Scanln(&resposta)
			if resposta == "s" {
				codigo := 0

				for {
					fmt.Println("Informe um nome para a marca: ")
					fmt.Scanln(&marca)
					marcas = append(marcas, marca)

					fmt.Println("Deseja cadastrar outra marca? s/n")
					fmt.Scanln(&resposta)

					if resposta != "s" {
						break
					}
				}

				rows, _ := db.Query("select max(marca_id) as ultimo from marcas")
				defer rows.Close()
				for rows.Next() {
					rows.Scan(&codigo)
					codigo = codigo + 1
				}

				totalregistros = len(marcas)
				tx, _ := db.Begin()
				stmt, _ := tx.Prepare("insert into marcas(marca_id, marca_nome) values(?,?)")

				for i := 0; i < totalregistros; i++ {
					stmt.Exec(codigo, marcas[i])
					codigo = codigo + 1
				}

				if err != nil {
					tx.Rollback()
					log.Fatal(err)
				}
				tx.Commit()
			}

		case "3":

			rows, _ := db.Query("select categoria_id, categoria_nome from categorias")
			defer rows.Close()
			fmt.Println("\nAs seguintes categorias estao registradas:")
			for rows.Next() {
				var c categoriaa
				rows.Scan(&c.id, &c.nome)
				fmt.Println(c)
			}

		case "4":

			rows, _ := db.Query("select marca_id, marca_nome from marcas")
			defer rows.Close()
			fmt.Println("\nAs seguintes marcas estao registradas:")
			for rows.Next() {
				var c marcaa
				rows.Scan(&c.id, &c.nome)
				fmt.Println(c)
			}

		case "5":
			fmt.Print("Informe o código da Categoria que deseja atualizar: ")
			fmt.Scanln(&codigo)

			fmt.Print("Informe o novo nome da Categoria que deseja atualizar: ")
			fmt.Scanln(&nome)

			// update
			stmt, _ := db.Prepare("Update categorias set categoria_nome = ? where categoria_id = ?")
			stmt.Exec(nome, codigo)

		case "6":
			fmt.Println("Informe o código da Marca que deseja atualizar: ")
			fmt.Scanln(&codigo)

			fmt.Println("Informe o novo nome da Marca que deseja atualizar: ")
			fmt.Scanln(&nome)

			// Update
			stmt, _ := db.Prepare("Update marcas set marca_nome = ? where marca_id = ?")
			stmt.Exec(nome, codigo)
		case "7":
			exec(db, `DROP TABLE categorias`)
			resposta = "0"

		case "8":
			exec(db, `DROP TABLE marcas`)
			resposta = "0"

		case "9":
			exec(db, `DROP TABLE categorias`)
			exec(db, `DROP TABLE marcas`)
			resposta = "0"

		case "0":
			break

		default:
			fmt.Println("Opção inválida.")
			break
		}

	}
}
