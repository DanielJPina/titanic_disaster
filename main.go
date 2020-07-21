package main

func main() {
	// Carregar ficheiros - chama função do ficheiro inputoutput.go
	registosTrain := carregarFicheiro("train.csv")
	registosTest := carregarFicheiro("test.csv")

	// Dados dos ficheiros csv sem cabeçalho
	trainInicial := registosTrain[1:]
	testInical := registosTest[1:]

	// Adicionar coluna de survived no test - chama função do ficheiro tratamentodados.go
	testInicialCompleto := adicionarColunaSurvived(testInical)

	// Tratamento de espaços em branco dos ficheiros - chama função do ficheiro tratamentodados.go
	alterarBranco(trainInicial)
	alterarBranco(testInicialCompleto)

	// Remoção de atributos não relevantes dos ficheiros - chama função do ficheiro tratamentodados.go
	train := eliminarAtributos(trainInicial)
	test := eliminarAtributos(testInicialCompleto)

	//Inicialização da arvore de decisão
	arvoreDecisao := &tree{
		root: nil,
	}

	// Algoritmo de criação de arvore de decisão - chama função do ficheiro arvoredecisao.go
	criacaoArvoreDecisao(arvoreDecisao, train)

	// Desenha e imprime a arvore de decisão - chama função do ficheiro desenhaarvore.go
	desenho := desenharArvore(arvoreDecisao.root)
	exportarTxt(desenho)

	// Analisar passageiros treino - chama função do ficheiro analise.go
	analisarTreino(train, arvoreDecisao)

	// Analisar passageiros teste - chama função do ficheiro analise.go
	output := analisarPassageiros(test, arvoreDecisao)

	// Converter output para ficheiro submission.csv - chama função do ficheiro inputoutput.go
	exportarCSV(output)
}
