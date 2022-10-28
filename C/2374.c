#include <stdio.h>
 
int main() {
    int bomba, pneu;
    scanf("%d", &bomba);
    scanf("%d", &pneu);
    int dif = bomba - pneu;
    printf("%d\n", dif);
    return 0;
}
