//矩阵
#include<stdio.h>

int main() {
		int i,j;
		int a,b;
		scanf("%d %d",&i,&j);
		double array[i][j];
		for(a=0;a<i;a++){
				for(b=0;b<j;b++){
						scanf("%lf",&array[a][b]);
				}
		}

		double tmp = array[0][0];
		printf("%f\n",tmp);
		int count=0;
		for(a=0;a<i;a++){
				for(b=0;b<j;b++){
						if(tmp>array[a][b]){
								tmp=array[a][b];
								count=0;
						}
						if(tmp==array[a][b])
								count++;
				}
		}
		printf("%d  %lf\n",count,tmp);
		return 0;
}
