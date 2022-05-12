func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
			select{
				case first := <-firstChan:
				output <- first * first
				case second := <-secondChan:
				output <- second * 3
				case <-stopChan:
				break
			}
	}()
	return output
}





