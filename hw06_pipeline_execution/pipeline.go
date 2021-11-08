package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	chIn := in

	for _, stage := range stages {
		chOut := make(Bi)
		stage := stage

		go func(chIn In) {
			defer close(chOut)
			chStage := stage(chIn)

			for {
				select {
				case val, ok := <-chStage:
					if !ok {
						return
					}
					chOut <- val
				case <-done:
					return
				}
			}
		}(chIn)

		chIn = chOut
	}

	return chIn
}
