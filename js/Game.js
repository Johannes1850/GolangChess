
class Game {
    constructor(width) {
        this.size = width;
        this.board = new Board(this.size*2);
        ctx.drawImage(img, 0, 0, this.size*2, this.size*2);
        this.board.Draw();
        this.nextMove = true
    }

    incrLoadingBar(data) {
        console.log(data);
        let bar = document.getElementById("bar");
        let progress = document.getElementById("progress");
        while (document.getElementById("bar").clientWidth < document.getElementById("progress").clientWidth*parseInt(data)/100) {
            let width = bar.clientWidth;
            let newWidth = ((width+1)/progress.offsetWidth)*100;
            bar.style.width=newWidth.toString()+"%";
        }
    }

    Draw() {
        this.board.Draw();
    }

    Update(posX, posY) {
        let tileX = posX / (this.size/8), tileY = posY / (this.size/8);
        tileX = Math.floor(tileX), tileY = Math.floor(tileY);
        this.board.Update(new Point(tileX, tileY));
        this.Draw();
        if (this.board.aiMove) {
            this.board.aiMove = false;
            this.board.blockHumanMove = true;
            $.ajax({
                url: 'receive',
                type: 'post',
                dataType: 'html',
                data : {next_move: false.toString(), board_position: this.board.position.dataFormat().toString()
                    , whiteKingMoved: this.board.position.whiteKingMoved.toString()
                    , blackKingMoved: this.board.position.blackKingMoved.toString()
                    , rookA1Moved: this.board.position.RookA1Moved.toString()
                    , rookH1Moved: this.board.position.RookH1Moved.toString()
                    , rookA8Moved: this.board.position.RookA8Moved.toString()
                    , rookH8Moved: this.board.position.RookH8Moved.toString()},
                success : this.board.PlayAiMove
            });
        }
    }
}