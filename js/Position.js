class Position {
    constructor() {
        this.blackPieces = [];
        this.whitePieces = [];
        this.whiteKingMoved = false;
        this.blackKingMoved = false;
        this.RookA1Moved = false;
        this.RookH1Moved = false;
        this.RookA8Moved = false;
        this.RookH8Moved = false;

    }

    clone() {
        let pos = new Position();
        for (let piece of this.whitePieces) {
            pos.whitePieces.push(piece.clone())
        }
        for (let piece of this.blackPieces) {
            pos.blackPieces.push(piece.clone())
        }
        return pos;
    }

    * allPieces() {
        for (let piece of this.blackPieces) {
            yield piece;
        }

        for (let piece of this.whitePieces) {
            yield piece;
        }
    }

    exposeKing(move) {
        let pos = this.clone();
        pos.PlayMoveAlways(move);
        let allMoves = [];
        for (let move of pos.allValidMoves("black", true)) {allMoves.push(move);}
        if (allMoves.length == 0) {return false}
        for (let move of allMoves) {
            let newPos = pos.clone()
            newPos.PlayMoveAlways(move)
            if (Math.abs(newPos.evaluation()) == 1) {return true}
        }
        return false
    }

    RemovePiece(piece) {
        if (piece.color == pieceColor.WHITE) {
            for (let i = 0; i < this.whitePieces.length; i++) {
                if (this.whitePieces[i].position == piece.position) {
                    this.whitePieces.splice(i, 1);
                    return;
                }
            }
        }

        if (piece.color == pieceColor.BLACK) {
            for (let i = 0; i < this.blackPieces.length; i++) {
                if (this.blackPieces[i].position == piece.position) {
                    this.blackPieces.splice(i, 1);
                    return;
                }
            }
        }
    }

    addQueen(pos) {
        if (pos.y == 0) {
            this.whitePieces.push(new Queen('white', pos));
        }
        if (pos.y == 7) {
            this.blackPieces.push(new Queen('black', pos));
        }
    }

    init() {
        this.blackPieces.push(new King('black', new Point(4, 0)));
        this.blackPieces.push(new Queen('black', new Point(3, 0)));
        this.blackPieces.push(new Knight('black', new Point(6, 0)));
        this.blackPieces.push(new Knight('black', new Point(1, 0)));
        this.blackPieces.push(new Rook('black', new Point(0, 0)));
        this.blackPieces.push(new Rook('black', new Point(7, 0)));
        this.blackPieces.push(new Bishop('black', new Point(5, 0)));
        this.blackPieces.push(new Bishop('black', new Point(2, 0)));
        this.blackPieces.push(new Pawn('black', new Point(0, 1)));
        this.blackPieces.push(new Pawn('black', new Point(1, 1)));
        this.blackPieces.push(new Pawn('black', new Point(2, 1)));
        this.blackPieces.push(new Pawn('black', new Point(3, 1)));
        this.blackPieces.push(new Pawn('black', new Point(4, 1)));
        this.blackPieces.push(new Pawn('black', new Point(5, 1)));
        this.blackPieces.push(new Pawn('black', new Point(6, 1)));
        this.blackPieces.push(new Pawn('black', new Point(7, 1)));

        this.whitePieces.push(new King('white', new Point(4, 7)));
        this.whitePieces.push(new Queen('white', new Point(3, 7)));
        this.whitePieces.push(new Knight('white', new Point(6, 7)));
        this.whitePieces.push(new Knight('white', new Point(1, 7)));
        this.whitePieces.push(new Rook('white', new Point(0, 7)));
        this.whitePieces.push(new Rook('white', new Point(7, 7)));
        this.whitePieces.push(new Bishop('white', new Point(5, 7)));
        this.whitePieces.push(new Bishop('white', new Point(2, 7)));
        this.whitePieces.push(new Pawn('white', new Point(0, 6)));
        this.whitePieces.push(new Pawn('white', new Point(1, 6)));
        this.whitePieces.push(new Pawn('white', new Point(2, 6)));
        this.whitePieces.push(new Pawn('white', new Point(3, 6)));
        this.whitePieces.push(new Pawn('white', new Point(4, 6)));
        this.whitePieces.push(new Pawn('white', new Point(5, 6)));
        this.whitePieces.push(new Pawn('white', new Point(6, 6)));
        this.whitePieces.push(new Pawn('white', new Point(7, 6)));
    }

    PlayMoveAlways(move) {
        let piece = this.PieceAt(move.start);
        this.RemovePiece(this.PieceAt(move.end));
        if (piece instanceof Rook) {
            if (piece.color === pieceColor.WHITE) {
                if (piece.position.x === 0) {this.RookA1Moved = true;}
                if (piece.position.x === 7) {this.RookH1Moved = true;}
            }
            if (piece.color === pieceColor.BLACK) {
                if (piece.position.x === 0) {this.RookA8Moved = true;}
                if (piece.position.x === 7) {this.RookH8Moved = true;}
            }
        }
        // castling
        if (piece instanceof King) {
            if (piece.color === pieceColor.WHITE && !this.whiteKingMoved) {
                this.whiteKingMoved = true;
                if (move.start.x-2 === move.end.x && !this.RookA1Moved) {
                    this.PieceAt({x:0,y:7}).updatePosition({x:3,y:7})
                    this.RookA1Moved = true;
                }
                if (move.start.x+2 === move.end.x && !this.RookH1Moved) {
                    this.PieceAt({x:7,y:7}).updatePosition({x:5,y:7})
                    this.RookH1Moved = true;
                }
            }
            if (piece.color === pieceColor.BLACK && !this.blackKingMoved) {
                this.blackKingMoved = true;
                if (move.start.x-2 === move.end.x && !this.RookA8Moved) {
                    this.PieceAt({x:0,y:0}).updatePosition({x:3,y:0})
                    this.RookA8Moved = true;
                }
                if (move.start.x+2 === move.end.x && !this.RookH8Moved) {
                    this.PieceAt({x:7,y:0}).updatePosition({x:5,y:0})
                    this.RookH8Moved = true;
                }
            }
        }
        // pawn promotion
        if (piece instanceof Pawn && (move.end.y == 7 || move.end.y == 0)) {
            this.RemovePiece(this.PieceAt(move.start));
            this.addQueen(move.end);
        } else {piece.updatePosition(move.end);}
    }

    PlayMove(move) {
        let piece = this.PieceAt(move.start);
        if (piece.isValidMove === "undefined") {return false;}
        if (piece.isValidMove(move.start, move.end, this)) {
            if (this.exposeKing(move)) {return false}
            if (this.PieceAtBool(move.end)) {
                this.RemovePiece(this.PieceAt(move.end));
            }
            // pawn promotion
            if (piece instanceof Pawn && (move.end.y == 7 || move.end.y == 0)) {
                this.RemovePiece(this.PieceAt(move.start));
                this.addQueen(move.end);
            }
            if (piece instanceof Rook) {
                if (piece.color === pieceColor.WHITE) {
                    if (piece.position.x === 0) {this.RookA1Moved = true;}
                    if (piece.position.x === 7) {this.RookH1Moved = true;}
                }
                if (piece.color === pieceColor.BLACK) {
                    if (piece.position.x === 0) {this.RookA8Moved = true;}
                    if (piece.position.x === 7) {this.RookH8Moved = true;}
                }
            }
            // castling
            if (piece instanceof King) {
                if (piece.color === pieceColor.WHITE && !this.whiteKingMoved) {
                    this.whiteKingMoved = true;
                    if (move.start.x-2 === move.end.x && !this.RookA1Moved) {
                        this.PieceAt({x:0,y:7}).updatePosition({x:3,y:7})
                        this.RookA1Moved = true;
                    }
                    if (move.start.x+2 === move.end.x && !this.RookH1Moved) {
                        this.PieceAt({x:7,y:7}).updatePosition({x:5,y:7})
                        this.RookH1Moved = true;
                    }
                }
                if (piece.color === pieceColor.BLACK && !this.blackKingMoved) {
                    this.blackKingMoved = true;
                    if (move.start.x-2 === move.end.x && !this.RookA8Moved) {
                        this.PieceAt({x:0,y:0}).updatePosition({x:3,y:0})
                        this.RookA8Moved = true;
                    }
                    if (move.start.x+2 === move.end.x && !this.RookH8Moved) {
                        this.PieceAt({x:7,y:0}).updatePosition({x:5,y:0})
                        this.RookH8Moved = true;
                    }
                }
            }

            if (piece instanceof King) {
                if (piece.color === pieceColor.WHITE && !this.whiteKingMoved) {
                    this.whiteKingMoved = true;
                }
                if (piece.color === pieceColor.BLACK && !this.blackKingMoved) {
                    this.blackKingMoved = true;
                }
            }

            if (piece instanceof Rook) {
                if (piece.color === pieceColor.WHITE) {
                    if (piece.position.x === 0 && piece.position.y === 7 && !this.RookA1Moved) {
                        this.RookA1Moved = true;
                    }
                    if (piece.position.x === 1 && piece.position.y === 1 && !this.RookH1Moved) {
                        this.RookH1Moved = true;
                    }
                }
                if (piece.color === pieceColor.BLACK) {
                    if (piece.position.x === 0 && piece.position.y === 0 && !this.RookA8Moved) {
                        this.RookA8Moved = true;
                    }
                    if (piece.position.x === 7 && piece.position.y === 0 && !this.RookH8Moved) {
                        this.RookH8Moved = true;
                    }
                }
            }
            piece.updatePosition(move.end);
            return true;
        } else {
            return false;
        }
    }

    PlayMoveReturnPos(move) {
        let pos = this.clone();
        pos.PieceAt(move.start).updatePosition(move.end);
        return pos.evaluation();
    }

    PieceAt(point) {
        for (let piece of this.allPieces()) {
            if (piece.position.x === point.x && piece.position.y === point.y) {
                return piece;
            }
        }
        return new EmptyPiece();
    }

    PieceAtBool(point) {
        for (let piece of this.allPieces()) {
            if (piece.position.x === point.x && piece.position.y === point.y) {
                return true;
            }
        }
        return false;
    }

    PieceAtBoolColor(point, color) {
        if (color === pieceColor.WHITE) {
            for (let piece of this.blackPieces) {
                if (piece.position.x === point.x && piece.position.y === point.y) {
                    return true;
                }
            }
        }
        if (color == pieceColor.BLACK) {
            for (let piece of this.whitePieces) {
                if (piece.position.x === point.x && piece.position.y === point.y) {
                    return true;
                }
            }
        }
        return false;
    }

    PieceAtBoolColor2(point, color) {
        if (color == pieceColor.WHITE) {
            for (let piece of this.whitePieces) {
                if (piece.position.x === point.x && piece.position.y === point.y) {
                    return true;
                }
            }
        }
        if (color == pieceColor.BLACK) {
            for (let piece of this.blackPieces) {
                if (piece.position.x === point.x && piece.position.y === point.y) {
                    return true;
                }
            }
        }
        return false;
    }

    blockingPiece(start, end, color) {
        if (color === pieceColor.WHITE) {
            for (let piece of this.whitePieces) {
                if (end.x === piece.position.x && end.y === piece.position.y) {
                    return true;
                }
            }
        }
        if (color === pieceColor.BLACK) {
            for (let piece of this.blackPieces) {
                if (end.x === piece.position.x && end.y === piece.position.y) {
                    return true;
                }
            }
        }

        // horizontal
        if (start.y == end.y && start.x != end.x) {
            let horizontalDiff = Math.abs(start.x-end.x);
            if (start.x < end.x) {
                for (let i = 1; i < horizontalDiff; i++) {
                    if (this.PieceAtBool(new Point(start.x+i, start.y))) { return true; }
                }
            }
            if (start.x > end.x) {
                for (let i = 1; i < horizontalDiff; i++) {
                    if (this.PieceAtBool(new Point(start.x-i, start.y))) { return true; }
                }
            }
        }

        // vertical
        if (start.y != end.y && start.x == end.x) {
            let verticalDiff = Math.abs(start.y-end.y);
            if (start.y < end.y) {
                for (let i = 1; i < verticalDiff; i++) {
                    if (this.PieceAtBool(new Point(start.x, start.y+i))) { return true; }
                }
            }
            if (start.y > end.y) {
                for (let i = 1; i < verticalDiff; i++) {
                    if (this.PieceAtBool(new Point(start.x, start.y-i))) { return true; }
                }
            }
        }

        //diagonal
        let diff = new Point(start.x-end.x, start.y-end.y);
        if (Math.abs(diff.x) == Math.abs(diff.y)) {
            // topLeft
            if (diff.x > 0 && diff.y > 0) {
                for (let i = 1; i < diff.x; i++) {
                    if (this.PieceAtBool(new Point(start.x-i, start.y-i))) { return true; }
                }
            }
            // topRight
            if (diff.x < 0 && diff.y > 0) {
                for (let i = 1; i < Math.abs(diff.x); i++) {
                    if (this.PieceAtBool(new Point(start.x+i, start.y-i))) { return true; }
                }
            }
            // bottomLeft
            if (diff.x > 0 && diff.y < 0) {
                for (let i = 1; i < Math.abs(diff.x); i++) {
                    if (this.PieceAtBool(new Point(start.x-i, start.y+i))) { return true; }
                }
            }
            //bottomRight
            if (diff.x < 0 && diff.y < 0) {
                for (let i = 1; i < Math.abs(diff.x); i++) {
                    if (this.PieceAtBool(new Point(start.x+i, start.y+i))) { return true; }
                }
            }
        }
        return false;
    }

    * allValidMoves(color, onlyHitting) {
        if (onlyHitting) {
            if (color === pieceColor.BLACK) {
                for (let piece of this.blackPieces) {
                    for (let move of piece.allValidMoves(this, color)) {
                        if (this.PieceAtBoolColor(move.end, color)) {
                            yield move;
                        }
                    }
                }
            }
            if (color === pieceColor.WHITE) {
                for (let piece of this.whitePieces) {
                    for (let move of piece.allValidMoves(this, color)) {
                        if (this.PieceAtBoolColor(move.end, color)) {
                            yield move;
                        }
                    }
                }
            }
        }
        else {
            if (color === pieceColor.BLACK) {
                for (let piece of this.blackPieces) {
                    for (let move of piece.allValidMoves(this, color)) {
                        yield move;
                    }
                }
                if (!this.blackKingMoved) {
                    if (!this.RookA8Moved) {
                        if (!this.blockingPiece(new Point(0,0), new Point(3,0), pieceColor.BLACK)) {
                            yield {start: new Point(0,4), end:new Point(0,2)};
                        }
                    }
                    if (!this.RookH8Moved) {
                        if (!this.blockingPiece(new Point(7,0), new Point(5,0), pieceColor.BLACK)) {
                            yield {start: new Point(0,4), end:new Point(0,6)};
                        }
                    }
                }
            }
            if (color === pieceColor.WHITE) {
                for (let piece of this.whitePieces) {
                    for (let move of piece.allValidMoves(this, color)) {
                        yield move;
                    }
                }
                if (!this.whiteKingMoved) {
                    if (!this.RookA1Moved) {
                        if (!this.blockingPiece(new Point(0,7), new Point(3,7), pieceColor.WHITE)) {
                            yield {start: new Point(7,4), end:new Point(7,2)};
                        }
                    }
                    if (!this.RookH1Moved) {
                        if (!this.blockingPiece(new Point(7,7), new Point(5,7), pieceColor.WHITE)) {
                            yield {start: new Point(7,4), end:new Point(7,6)};
                        }
                    }
                }
            }
        }
    }

    allValidMovesSorted(color) {
        let allMoves = [];
        for (let move of this.allValidMoves(color)) {allMoves.push(move)}
        if (color === pieceColor.BLACK) {
            for (let move of allMoves) {
                // console.log(move.start.x+" "+move.end.x);
            }
        }
        let firstMove = allMoves[0];
        if (color == pieceColor.BLACK) {
            allMoves.sort((a, b) => this.PieceAt(b.end).value-this.PieceAt(a.end).value);
        }
        if (color === pieceColor.WHITE) {
            allMoves.sort((a, b) => this.PieceAt(b.end).value-this.PieceAt(a.end).value);
        }
        if (color === pieceColor.BLACK) {
            for (let move of allMoves) {
                // console.log(move.start.x+" "+move.end.x);
            }
        }
        // console.log(ChangedOrder);
        return allMoves;
    }

    Draw() {
        for (let piece of this.allPieces()) {
            piece.Draw();
        }
    }

    evaluation() {
        let whiteCount = 0;
        let blackCount = 0;
        let blackKingExists = false;
        let whiteKingExists = false;
        for (let piece of this.allPieces()) {
            if (piece.color === pieceColor.WHITE) {
                if (piece instanceof King) {whiteKingExists = true}
                whiteCount += piece.value;
            }
            if (piece.color === pieceColor.BLACK) {
                if (piece instanceof King) {blackKingExists = true}
                blackCount += piece.value;
            }
        }
        if (!blackKingExists) {return 1;}
        if (!whiteKingExists) {return -1;}
        return (whiteCount/blackCount-1);
    }

    dataFormat() {
        let finalArray = new Float32Array(768);
        for (let piece of this.allPieces()) {
            if (piece.color == pieceColor.WHITE) {
                if (piece instanceof Pawn) {
                    finalArray[this.pointToChessTile(piece.position)] = 1.0;
                }
                if (piece instanceof King) {
                    finalArray[this.pointToChessTile(piece.position)+64] = 1.0;
                }
                if (piece instanceof Queen) {
                    finalArray[this.pointToChessTile(piece.position)+64*2] = 1.0;
                }
                if (piece instanceof Bishop) {
                    finalArray[this.pointToChessTile(piece.position)+64*3] = 1.0;
                }
                if (piece instanceof Rook) {
                    finalArray[this.pointToChessTile(piece.position)+64*4] = 1.0;
                }
                if (piece instanceof Knight) {
                    finalArray[this.pointToChessTile(piece.position)+64*5] = 1.0;
                }
            }
            if (piece.color == pieceColor.BLACK) {
                if (piece instanceof Pawn) {
                    finalArray[this.pointToChessTile(piece.position)+64*6] = 1;
                }
                if (piece instanceof King) {
                    finalArray[this.pointToChessTile(piece.position)+64*7] = 1;
                }
                if (piece instanceof Queen) {
                    finalArray[this.pointToChessTile(piece.position)+64*8] = 1;
                }
                if (piece instanceof Bishop) {
                    finalArray[this.pointToChessTile(piece.position)+64*9] = 1;
                }
                if (piece instanceof Rook) {
                    finalArray[this.pointToChessTile(piece.position)+64*10] = 1;
                }
                if (piece instanceof Knight) {
                    finalArray[this.pointToChessTile(piece.position)+64*11] = 1;
                }
            }
        }
        return finalArray;
    }

    pointToChessTile(point) {
        let posX = point.x+1;
        let posY = 7-point.y;
        return ((posY*8)+posX);
    }
}