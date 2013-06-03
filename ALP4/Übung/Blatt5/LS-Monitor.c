#include <pthread.h>

//Loesung analog zu der in Go

//Kompiliert so nicht da keine main(), sonst keine Errors

unsigned int nR = 0; 
unsigned int nW = 0;

unsigned int rCnt = 0;
unsigned int wCnt = 0;

pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

pthread_cond_t okR = PTHREAD_COND_INITIALIZER;
pthread_cond_t okW = PTHREAD_COND_INITIALIZER;

void readerIn(){
	pthread_mutex_lock(&mutex);

	if (nW > 0 || wCnt>0) {
		rCnt++;	
		pthread_cond_wait(&okR, &mutex); 
		rCnt--;
	}
	nR++;

	pthread_mutex_unlock(&mutex);
}

void readerOut(){
	pthread_mutex_lock(&mutex);
	nR--;
	if (nR == 0) {
		pthread_cond_signal(&okW);
	}
	pthread_mutex_unlock(&mutex);
}

void writerIn(){
	pthread_mutex_lock(&mutex);
	if (nR > 0 || nW > 0) {
		wCnt++;
		pthread_cond_wait(&okW, &mutex); 
		wCnt--;
	}
	nW = 1;
	pthread_mutex_unlock(&mutex);
}

void writerOut(){
	pthread_mutex_lock(&mutex);
	nW = 0;

	if (rCnt>0) {
		pthread_cond_signal(&okR);
	} else {
		pthread_cond_signal(&okW);
	}
	pthread_mutex_unlock(&mutex);
}