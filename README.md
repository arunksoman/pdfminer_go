# pdfminer_go

There is no package entirely written in Go that can extract text from PDF. The packages like fitz, PDFBox which is not written in go. So here I'm trying to write a package in pure Go without c dependency. This is a port of Python [pdfminer.six](https://github.com/pdfminer/pdfminer.six) which is community maintained fork of [pdfminer](https://github.com/euske/pdfminer). Referenced also [https://github.com/dslipak/pdf](https://github.com/dslipak/pdf)

#### Its under development. Trying to make an alpha release after 6 months today (26th May 2024).

Todos

- [ ] Added RC4 and Test
- [ ] Add ASCII85 Decode --> Adobe Implementation
- [ ] Add Fonts
- [X] Add ASCII85 HexDecode
- [ ] Support For different PDF compression
- [ ] Support till PDF 1.7 standard
- [ ] Support For PDF 2.0 standard
- [ ] Text extraction algorithm

Problems

- [X] I know now its not a proper GO package structure. I am not a golang developer. As I gain more experience and everything falls in place, I will restructure the project. 
- [X] Now I am not thinking too much about code quality. I will rethink about that when things falls everything in place.
