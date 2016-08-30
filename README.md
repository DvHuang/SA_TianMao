# SA_TianMao
Sentiment Analysis for commodity command on TianMao


**About original:**

We comapre three way to ayalysis it 

1. basic Sentiment word. 
2. Doc2vec, use the doc2vec[3]'s way to classify the command,then classify the new command.
3. sentence autodecoder,in fact it was a seq2seq model,in the paper[4],we commend the Y=X,and it from  supervised to unsupervised ,then the encoder could be a feature extract like autocode in CNN.

**Papers:**

* [ [1] Sequence to Sequence Learning with Neural Networks](http://papers.nips.cc/paper/5346-sequence-to-sequence-learning-with-neural-networks.pdf)
* [ [2] Semi supervised Sequence Learning](https://arxiv.org/abs/1511.01432)
* [ [3] Distributed Representations of Sentences and Documents](https://arxiv.org/abs/1405.4053)
* [ [4] Skip thought vectors](http://arxiv.org/abs/1506.06726)



**File abstract:**

  there were the first approach in here,if you wnat the two and three way please connect me


  
## License

[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)
