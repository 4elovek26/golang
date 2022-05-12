func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)
	sum := 0

	go func() {
		defer close(output)
		for {
			select{
				case value := <-arguments:
					sum += value
				case <-done:
					output <- sum
					return
			}
		}
	}()
	return output
}



