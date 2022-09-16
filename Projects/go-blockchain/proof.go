package blockchain

//proof of work algorithm

//Take the data from the block

//create a counter (nonce) which starts at 0

//create a hash of the ata plus the counter

//check the hash to see if it meets a set of requirements

//Requirements:
//The first few bytes must contains 0s

const Difficulty = 12 

type ProofOfWork strunct {
	Block *Block
	Target *big.Int

}


fuc NewProof(b *Block) *ProofOfWork {
	traget := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int)

