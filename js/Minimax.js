class Minimax {
    constructor() {
        this.nextMove = {eval: 0, move: {start: new Point(0,0), end: new Point(0,0)}, third: true};
        this.allDepth4Moves = [];
        this.starttime = performance.now();
        this.position = undefined;
        this.MAX_LEVEL = 0;
    }

    nextAiMove(position, color) {
        this.nextMove = {eval: 0, move: {start: new Point(0,0), end: new Point(0,0)}, third: true};
        this.starttime = performance.now();
        this.position = position;
        this.allDepth4Moves = [];
        this.MAX_LEVEL = 4;
        let minCalculation = this.minimax(position, 1, Number.NEGATIVE_INFINITY, Number.POSITIVE_INFINITY, color
            , {start: new Point(0,0), end: new Point(0,0)}, false);
        let avgEval = 0;
        this.allDepth4Moves.sort((a, b) => b.eval-a.eval);
        this.allDepth4Moves = this.allDepth4Moves.slice(0,2);
        this.MAX_LEVEL = 6;
        //let minCalculation2 = this.minimax(position, 1, Number.NEGATIVE_INFINITY, Number.POSITIVE_INFINITY, color
        //    , {start: new Point(0,0), end: new Point(0,0)}, true);

        let para = document.createElement("p");
        let t = document.createTextNode("("+this.nextMove.move.start.x.toString() + ","
            + this.nextMove.move.start.y.toString() + ")(" + this.nextMove.move.end.x.toString() + ","
            + this.nextMove.move.end.y.toString() + ")"+" "+this.nextMove.eval.toString());
        para.appendChild(t);
        document.getElementById("gameInfo").appendChild(para);
        position.PlayMove(this.nextMove.move);
    }

    minimax(position, depth, alpha, beta, color, prevMove, secondRun) {
        let posEval = position.evaluation();
        if (depth >= this.MAX_LEVEL || Math.abs(posEval) === 1) {
            return {eval: posEval, prevMove: prevMove};
        }
        let date = performance.now();
        // return posEval at MAX_LEVEL
        let allValidMoves = [];
        if (!secondRun || depth > 1) {
            if (depth >= this.MAX_LEVEL) {
                for (let move of position.allValidMoves(color, true)) {allValidMoves.push(move);}
            } else {
                allValidMoves = position.allValidMovesSorted(color);
            }
        } else {
            for (let move of this.allDepth4Moves) {allValidMoves.push(move.move);}
            // for (let move of allValidMoves) {console.log(move.start);}
        }

        if ((depth >= this.MAX_LEVEL && allValidMoves.length === 0) || date/1000-this.starttime/1000 > 50) {
            return {eval: posEval, prevMove: prevMove};
        }

        if (color == pieceColor.BLACK) {
            let minEval = {eval: Number.POSITIVE_INFINITY, prevMove: {start: new Point(0,0), end: new Point(0,0)}};
            for (let move of allValidMoves) {
                let pos = position.clone();
                pos.PlayMove(move);

                // recursive call
                let a = this.minimax(pos, depth + 1, alpha, beta, pieceColor.WHITE, move, secondRun);

                if (a.eval < minEval.eval) minEval = a;
                beta = Math.min(beta, a.eval);
                if (beta <= alpha) break;
            }

            // Adds all solutions at depth 2 to list
            if (depth === 2)
            {
                if (!secondRun) {
                    this.allDepth4Moves.push({eval: minEval.eval, move: prevMove});
                }
                if (minEval.eval > this.nextMove.eval || this.nextMove.third)
                {
                    this.nextMove = {eval: minEval.eval, move: prevMove, third: false};
                }
            }
            return minEval;
        }
        if (color === pieceColor.WHITE) {
            let maxEval = {eval: Number.NEGATIVE_INFINITY, prevMove: {start: new Point(0,0), end: new Point(0,0)}};
            for (let move of allValidMoves) {
                let pos = position.clone();
                pos.PlayMove(move);

                // recursive call
                let a = this.minimax(pos, depth + 1, alpha, beta, pieceColor.BLACK, move, secondRun);

                if (a.eval > maxEval.eval) maxEval = a;
                alpha = Math.max(alpha, a.eval);
                if (beta <= alpha) break;
            }
            // Adds all solutions at depth 2 to list
            if (depth === 2)
            {
                if (!secondRun) {
                    this.allDepth4Moves.push({eval: maxEval.eval, move: prevMove});
                }
                if (maxEval.eval < this.nextMove.eval || this.nextMove.third)
                {
                    this.nextMove = {eval: maxEval.eval, move: prevMove, third: false};
                }
            }
            return maxEval;
        }
    }
}
