class Board {
    constructor(size) {
        this.position = new Position();
        this.position.init();
        this.pieceSelected = false;
        this.selectedPiece = null;
        this.AiPlayer = new Minimax();
        this.aiMove = false;
        this.size = size;
        this.nextAiMove = ""
    }

    PlayAiMove(move) {
        console.log(move);
        MainGame.board.position.PlayMoveAlways({start: new Point(0, 1), end: new Point(0,3)});
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
            if (!this.position.PlayMove({start: this.selectedPiece.position, end: ClickedAt})) {
                this.pieceSelected = false;
                this.selectedPiece.deSelect();
                return;
            } else {
                this.pieceSelected = false;
                this.selectedPiece.deSelect();
                this.aiMove = true;
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