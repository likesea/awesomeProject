package main

import "strings"
var result [][]string = make([][]string,0)
func solveNQueens(n int) [][]string {
	board:=initBoard(n)
	backtrace(0,board)
	return result
}
func isValid(row int,col int,board []string) bool{
	 n := len(board)
	 //检查列
	 for i:=0;i<n;i++{
	 	if board[i][col]=='Q'{
	 		return false
		}
	 }
	 //检查右上方
	 for i,j:= row-1,col+1;i>=0&&j<n;i, j= i-1,j+1{
	 	if board[i][j]=='Q'{
	 		return false
		}
	}
	//左上方
	for i,j:=row-1,col-1;i>=0&&j>=0;i,j=i-1,j-1{
		if board[i][j]=='Q'{
			return false
		}
	}
	return true
}
  func backtrace (row int,board []string) {
	if row==len(board){
		cp :=make([]string,row)
		copy(cp,board)
		result = append(result, cp)
		return
	}
	n:=len(board[row])
	for col:=0;col<n;col++{
		if isValid(row ,col ,board){
			rowByte := []byte (board[row])
			rowByte[col]='Q'
			board[row] = string(rowByte)
			backtrace(row+1,board)
			rowByte = []byte (board[row])
			rowByte[col]='.'
			board[row] = string(rowByte)
		}else{
			continue
		}
	}
}


func initBoard(n int)[]string{
	board := make([]string,n)
	for i:=0;i<n;i++{
		board[i]=strings.Repeat(".", n)
	}
	return board
}
