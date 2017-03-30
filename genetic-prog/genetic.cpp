#include <cstdio>
#include <cstdlib>
#include <ctime>
#include <vector>

using namespace std;

class Chromosome {
  public:
    double heightMult, lineMult, holeMult;
};

int main() {
  srand(time(NULL));

  // Start with 16 random chromosomes.
  FILE* fp = fopen("init-pop", "r");
  vector<Chromosome*> population;
  while (!feof(fp)) {
    Chromosome* newChromo = new Chromosome;
    fscanf(fp,"%lf %lf %lf\n", &(newChromo->heightMult), &(newChromo->lineMult), &(newChromo->holeMult));

    population.push_back(newChromo);
  }
  fclose(fp);

  while (1) {
    // Select two chromos and make them compete.
    // Do it for all 16 chromos.

    // Don't really need to store the losers here, but I might change the algo
    // to keep a few losers and use them to crossover with the winners.
    vector<Chromosome*> winners,losers;
    while(!population.empty()) {
      int first = rand()%population.size();
      int second = rand()%population.size();
      if (second == first) {
        if (first>0) {
          second = first-1;
        } else {
          second = first+1;
        }
      }

      int winner, loser;

      // Run the first one 100 times and note the score.
      // Run the second one 100 times and note the score.

      // Set winner and loser accordingly.

      winners.push_back(population[winner]);
      losers.push_back(population[loser]);

      population.erase(population.begin()+winner);
      if (winner<loser) {
        population.erase(population.begin()+loser-1);
      } else {
        population.erase(population.begin()+loser);
      }
    }

    // We will discard the losers and crossover the winners.
    // Need to add mutations (with prob=10%) too!
    delete losers;
    for (int i=0;i<8;++i) {
      Chromosome* newChromo = new Chromosome;
      int r1 = rand()%2;
      int r2 = rand()%2;
      int r3 = rand()%2;

      if (r1==0) {
        newChromo->heightMult = winners[i]->heightMult;
      } else {
        newChromo->heightMult = winners[i+8]->heightMult;
      }

      if (r2==0) {
        newChromo->lineMult = winners[i]->lineMult;
      } else {
        newChromo->lineMult = winners[i+8]->lineMult;
      }

      if (r3==0) {
        newChromo->holeMult = winners[i]->holeMult;
      } else {
        newChromo->holeMult = winners[i+8]->holeMult;
      }

      population.push_back(newChromo);
    }

    for (int i=0; i<8; ++i) {
      population.push_back(winners[i]);
    }

    delete winners;
  }

  return 0;
}
