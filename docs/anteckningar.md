feature konvention för issues - "feat/lägg till en feature"
committa INTE till main

git branch för att se alla lokala branches, * visar vilken branch man just nu är i
"git checkout main" visar innehåll i main 
"git checkout -b feat/#8/open-empty-window" skapar ny branch med givet namn och "går till den"
kan inte bli för mycket branches, skapa ny branch för varje issue ex
"git diff" visar ändring mot senaste commit
"git log" visar historik av commits


committa helst inte med -m, använd ist. text editor och skriv ett ordentligt meddelande.

<Kort sammanfattning>

<Utförligare beskrivning>

Blankrad mellan gör att endast det på första raden (kort sammanfattning) visar i historiken på GitHub

"feat/8/open-empty-window has no upstream branch" felmeddeleande pga finns ingen branch på GitHub, endast lokalt.
"git push -u origin feat/8/open-empty-window" pushar och skapar en ny branch på GitHub

Hur får man in ändringar i main? Man gör en pull request, resultat av detta är en merge. Öppna en ny pull request på GitHub när en issue är closed. "Closes #8" som en beskrivning i pull requests, detta closar branchen automatiskt när den mergar."

Lös en merge-konflikt
1) checka ut main 2) kör git pull 3) checka ut konflikt-branchen 4) git merge main 5) Gå in i alla berörda filer och fixa
6) git commit 7) git push