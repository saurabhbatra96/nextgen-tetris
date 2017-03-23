#include<cstdio>

int main() {
  FILE* fp = fopen("greedy", "r");
  long int buffer,sum=0,num=0;

  while(!feof(fp)) {
    fscanf(fp,"%ld",&buffer);
    sum+=buffer;
    num++;
  }

  fclose(fp);
  printf("Average:%ld Number:%ld Sum: %ld",sum/num, num, sum);


  return 0;
}
