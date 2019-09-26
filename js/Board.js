class Board {
    constructor(size) {
        this.position = new Position();
        this.position.init();
        this.pieceSelected = false;
        this.selectedPiece = null;
        this.aiMove = false;
        this.size = size;
        this.blockHumanMove = false;
        this.nextAiMove = ""
    }

    PlayAiMove(move) {
        let formMove = move.split(',');
        MainGame.board.position.PlayMoveAlways({start: new Point(parseInt(formMove[0]
                , 10), parseInt(formMove[1], 10))
            , end: new Point(parseInt(formMove[2], 10),parseInt(formMove[3], 10))});
        MainGame.board.blockHumanMove = false;
        $.ajax({
            url: 'deepEval',
            type: 'post',
            success: function (data) {
                adjustEvalBar(data);
            }
        });
        MainGame.board.Draw();
    }

    Draw() {
        ctx.drawImage(img, 0, 0, this.size, this.size);
        this.position.Draw();
    }

    Update(ClickedAt = new Point(-1,-1)) {
        /**
        if (this.aiMove && ClickedAt.x === -1) {
            this.AiPlayer.nextAiMove(this.position, pieceColor.BLACK);
            this.aiMove = false;
            return;
        }
         **/
        if (this.pieceSelected) {
            if (this.position.PieceAtBool(ClickedAt) && this.position.PieceAt(ClickedAt).color === this.selectedPiece.color) {
                this.selectedPiece.deSelect();
                this.selectedPiece = this.position.PieceAt(ClickedAt);
                this.selectedPiece.select();
                return;
            }
            if (this.blockHumanMove) { return; }
            if (!this.position.PlayMove({start: this.selectedPiece.position, end: ClickedAt})) {
                this.pieceSelected = false;
                this.selectedPiece.deSelect();
                return;
            } else {
                document.getElementById("bar").style.width="0%";
                this.pieceSelected = false;
                this.selectedPiece.deSelect();
                this.aiMove = true;
                this.nextMove = false;
                this.pieceSelected = false;
                this.Draw();
                this.Update();
                return;
            }
        }

        if (!this.pieceSelected) {
            if (this.position.PieceAtBool(ClickedAt) && this.aiMove === false) {
                this.selectedPiece = this.position.PieceAt(ClickedAt);
                this.selectedPiece.select();
                this.pieceSelected = true;
            }
        }
    }
}