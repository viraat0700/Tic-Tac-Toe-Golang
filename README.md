# Tic-Tac-Toe-Golang
Lets connect local repo to git repo "Latest Test"

1. initBoard(): Pehla step hai ki ek function banaye, jisme game ka board shuruaati khali jagahon se bhara jaye. Yeh basically game ka shuruaati setup hai.

2. printBoard(): Doosra step hai ki ek function banaye, jo game board ka current state terminal par print kare. Aisa karke hum dekh payenge ki board kis tarah ka dikhta hai aur kisi bhi aur function ko debug karne mein madad milegi.

3. isValidMove(row, col int) bool: Tisra step hai ki ek function banaye, jo check kare ki player ka move sahi hai ya nahi. Matlab, yeh dekhega ki player ne sahi jagah par move kiya hai ya nahi, aur woh jagah pehle se occupied toh nahi hai.

4. checkWin() bool: Ab ek function banaye, jo dekhe ki kya koi player game jeet gaya hai. Yeh function board par sabhi possible winning combinations ko check karega.

5. checkDraw() bool: Fir ek function banaye, jo dekhe ki game draw ho gaya hai ya nahi. Yeh function board par koi khali jagah bach gayi hai ya sab bhar gaye hain, yeh dekhega.

6. switchPlayer(): Phir ek function banaye, jo har move ke baad players ke beech switch kare. Yani, 'X' aur 'O' players ke beech badlav kare.

7. main(): Ant mein, ek mukhya game loop implement karein jahan par players apne apne moves lete hain jab tak koi winner nahi hota ya game draw nahi ho jata.


YOU CAN RUN THE FILE USING ---> go run game.go