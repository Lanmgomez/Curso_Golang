// Em Go, os erros nao sao tradados em blocos de try catch

// o que acontece eh que na linha de chamada da funçao, ele mesmo me retorna esse erro
// e declaramos uma variavel para receber esse erro

// exemplo:
// res, err := http.Get("http://google.com")

// onde a variavel 'res' vai receber a resposta positiva 
// e a variavel 'err' vai receber os erros