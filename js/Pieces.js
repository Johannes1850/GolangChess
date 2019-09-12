const pieceColor = {
    WHITE: 'white',
    BLACK: 'black'
};

function Point(x, y) {
    this.x = x;
    this.y = y;
}

class Piece {
    constructor(value, position, color) {
        this.color = color;
        this.value = value;
        this.sprite = new Image();
        this.position = position;
        this.selected = false;
    }

    Draw() {
        if (this.selected) {
            ctx.drawImage(this.sprite, this.position.x * 152 + 5, this.position.y * 152 + 5, 120, 120);
            return;
        }
        ctx.drawImage(this.sprite, this.position.x * 152 + 15, this.position.y * 152 + 15, 100, 100);
    }

    updatePosition(newPos) {
        this.position = newPos;
    }

    select() {
        this.selected = true;
    }

    deSelect() {
        this.selected = false;
    }
}

class Pawn extends Piece {
    constructor(color, position) {
        super(1, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/PawnWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/PawnBlack.png";
                break;
        }
    }

    clone() {
        return new Pawn(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        if (this.color == pieceColor.WHITE) {
            // diagonal taking
            if (start.y-1 == end.y && (start.x-1 == end.x || start.x+1 == end.x) && position.PieceAtBoolColor(end, this.color)) {
                return true;
            }
            if (position.blockingPiece(start, end, 'white') || position.PieceAtBool(end)) { return false; }
            // hasn't been moved yet
            if (start.y == 6) {
                if ((start.y - 1 == end.y || start.y - 2 == end.y) && start.x == end.x) {
                    return true;
                }
            } else {
                if (start.y - 1 == end.y && start.x == end.x) {
                    return true;
                }
            }
        }

        if (this.color == pieceColor.BLACK) {
            // diagonal taking
            if (start.y+1 == end.y && (start.x-1 == end.x || start.x+1 == end.x) && position.PieceAtBoolColor(end, this.color)) {
                return true;
            }
            if (position.blockingPiece(start, end, 'black') || position.PieceAtBool(end)) { return false; }
            // hasn't been moved yet
            if (start.y == 1) {
                if ((start.y + 1 == end.y || start.y + 2 == end.y) && start.x == end.x) {
                    return true;
                }
            } else {
                if (start.y + 1 == end.y && start.x == end.x) {
                    return true;
                }
            }
        }
    }

    * allValidMoves(boardPosition, color) {
        if (color == pieceColor.WHITE) {
            if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y-1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-1, this.position.y-1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y-1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+1, this.position.y-1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, this.position.y-1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, this.position.y-1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, this.position.y-2), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, this.position.y-2)};
            }
        }
        if (color == pieceColor.BLACK) {
            if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y+1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-1, this.position.y+1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y+1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+1, this.position.y+1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, this.position.y+1), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, this.position.y+1)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, this.position.y+2), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, this.position.y+2)};
            }
        }
    }
}

class Rook extends Piece {
    constructor(color, position) {
        super(5, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/RookWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/RookBlack.png";
                break;
        }
    }

    clone() {
        return new Rook(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        if (position.blockingPiece(start, end, this.color)) { return false; }
        if (start.x != end.x && start.y == end.y || start.y != end.y && start.x == end.x) return true;
    }

    * allValidMoves(boardPosition, color) {
        for (let offset = 0; offset <= 7; offset++) {
            if (this.isValidMove(this.position, new Point(offset, this.position.y), boardPosition)) {
                yield {start: this.position, end: new Point(offset, this.position.y)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, offset)};
            }
        }
    }
}

class Bishop extends Piece {
    constructor(color, position) {
        super(3, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/BishopWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/BishopBlack.png";
                break;
        }
    }

    clone() {
        return new Bishop(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        if (position.blockingPiece(start, end, this.color)) { return false; }
        let diff = new Point(start.x-end.x, start.y-end.y);
        if (Math.abs(diff.x) == Math.abs(diff.y)) return true;
    }

    * allValidMoves(boardPosition, color) {
        for (let offset = 0; offset <= 7; offset++) {
            if (this.isValidMove(this.position, new Point(this.position.x+offset, this.position.y+offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+offset, this.position.y+offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x+offset, this.position.y-offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+offset, this.position.y-offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x-offset, this.position.y+offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-offset, this.position.y+offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x-offset, this.position.y-offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-offset, this.position.y-offset)};
            }
        }
    }
}

class Knight extends Piece {
    constructor(color, position) {
        super(3, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/KnightWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/KnightBlack.png";
                break;
        }
    }

    clone() {
        return new Knight(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.blockingPiece(start, end, this.color)) { return false; }
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        let diff = new Point(start.x-end.x, start.y-end.y);
        if ((Math.abs(diff.x) == 2 && Math.abs(diff.y) == 1)
            || (Math.abs(diff.x) == 1 && Math.abs(diff.y) == 2)) return true;
    }

    * allValidMoves(boardPosition, color) {
        if (this.isValidMove(this.position, new Point(this.position.x+2, this.position.y-1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+2, this.position.y-1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x+2, this.position.y+1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+2, this.position.y+1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y-2), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+1, this.position.y-2)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y+2), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+1, this.position.y+2)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-2, this.position.y+1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-2, this.position.y+1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-2, this.position.y-1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-2, this.position.y-1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y+2), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-1, this.position.y+2)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y-2), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-1, this.position.y-2)};
        }
    }
}

class Queen extends Piece {
    constructor(color, position) {
        super(9, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/QueenWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/QueenBlack.png";
                break;
        }
    }

    clone() {
        return new Queen(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        if (position.blockingPiece(start, end, this.color)) { return false; }
        let diff = new Point(start.x-end.x, start.y-end.y);
        if (Math.abs(diff.x) == Math.abs(diff.y)
            || (start.x != end.x && start.y == end.y || start.y != end.y && start.x == end.x)) return true;
    }

    * allValidMoves(boardPosition, color) {
        for (let offset = 0; offset <= 7; offset++) {
            if (this.isValidMove(this.position, new Point(offset, this.position.y), boardPosition)) {
                yield {start: this.position, end: new Point(offset, this.position.y)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x, offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x, offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x+offset, this.position.y+offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+offset, this.position.y+offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x+offset, this.position.y-offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x+offset, this.position.y-offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x-offset, this.position.y+offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-offset, this.position.y+offset)};
            }
            if (this.isValidMove(this.position, new Point(this.position.x-offset, this.position.y-offset), boardPosition)) {
                yield {start: this.position, end: new Point(this.position.x-offset, this.position.y-offset)};
            }
        }
    }
}

class King extends Piece {
    constructor(color, position) {
        super(10, position, color);
        switch (color) {
            case pieceColor.WHITE:
                this.sprite.src = "images/KingWhite.png";
                break;
            case pieceColor.BLACK:
                this.sprite.src = "images/KingBlack.png";
                break;
        }
    }

    clone() {
        return new King(this.color, new Point(this.position.x, this.position.y));
    }

    isValidMove(start, end, position) {
        if (position.PieceAtBoolColor2(end, this.color) || end.x < 0 || end.y < 0 || end.x > 7 || end.y > 7) {
            return false;
        }
        if (position.blockingPiece(start, end, this.color)) { return false; }
        let diff = new Point(start.x-end.x, start.y-end.y);
        if (Math.abs(diff.x) <= 1 && Math.abs(diff.y) <= 1) return true;
    }

    * allValidMoves(boardPosition, color) {
        if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+1, this.position.y)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y+1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+1, this.position.y+1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x+1, this.position.y-1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x+1, this.position.y-1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-1, this.position.y)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y+1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-1, this.position.y+1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x-1, this.position.y-1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x-1, this.position.y-1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x, this.position.y+1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x, this.position.y+1)};
        }
        if (this.isValidMove(this.position, new Point(this.position.x, this.position.y-1), boardPosition)) {
            yield {start: this.position, end: new Point(this.position.x, this.position.y-1)};
        }
    }
}

class EmptyPiece {
    constructor() {
        this.value = 0;
    }
}