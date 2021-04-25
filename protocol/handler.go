package protocol

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/GopherConRu/pb-fuzz-workshop/kv"
)

type handler struct {
	kv kv.KV
}

func NewHandler(kv kv.KV) *handler {
	return &handler{
		kv: kv,
	}
}

func (h *handler) NewConn() (io.WriteCloser, *bufio.Reader) {
	inputP, input := io.Pipe()
	output, outputP := io.Pipe()

	go h.run(bufio.NewReader(inputP), outputP)

	return input, bufio.NewReader(output)
}

func (h *handler) run(input *bufio.Reader, output io.WriteCloser) {
	defer output.Close()

	for {
		line, err := input.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("run: %s", err)
			}
			return
		}

		parts := strings.Split(line, " ")
		switch parts[0] {
		case "PING":
			fmt.Fprint(output, "+PONG\n")

		case "GET":
			o, err := h.kv.Get(kv.Key(parts[1]))
			if err != nil {
				if err == kv.ErrNotFound {
					fmt.Fprint(output, "$-1\n")
					continue
				}

				fmt.Fprintf(output, "-ERR %s.\n", err)
				continue
			}

			fmt.Fprintf(output, "$%d\n%s\n", len(o.Value), o.Value)

		case "SET":
			o, err := h.kv.Get(kv.Key(parts[1]))
			if err == kv.ErrNotFound {
				err = nil
			}
			if err != nil {
				fmt.Fprintf(output, "-ERR %s.\n", err)
				continue
			}

			if o == nil {
				o = new(kv.Object)
			}

			o.Value = []byte(parts[2])
			if err = h.kv.Set(kv.Key(parts[1]), o); err != nil {
				fmt.Fprintf(output, "-ERR %s.\n", err)
				continue
			}

			fmt.Fprintf(output, "+OK\n")
		}
	}
}
