import React from 'react'
import {PaginationInfo} from "../../types/tasks";
import ReactPaginate from "react-paginate";
import styled from 'styled-components';
import MTable from '@mui/material/Table';
import Typography from '@mui/material/Typography';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TablePagination from '@mui/material/TablePagination';
import Paper from '@mui/material/Paper';
import Container from "@mui/material/Container";

// const Pagination = styled(ReactPaginate).attrs({
//     // You can redifine classes here, if you want.
//     activeClassName: 'active', // default to "disabled"
// })`
//   margin-bottom: 2rem;
//   display: flex;
//   flex-direction: row;
//   justify-content: space-between;
//   list-style-type: none;
//   padding: 0 5rem;
//
//   li a {
//     border-radius: 7px;
//     padding: 0.1rem 1rem;
//     border: gray 1px solid;
//     cursor: pointer;
//   }
//
//   li.previous a,
//   li.next a,
//   li.break a {
//     border-color: transparent;
//   }
//
//   li.active a {
//     background-color: #0366d6;
//     border-color: transparent;
//     color: white;
//     min-width: 32px;
//   }
//
//   li.disabled a {
//     color: grey;
//   }
//
//   li.disable,
//   li.disabled a {
//     cursor: default;
//   }
// `;
//

export interface OnPageClickParams {
    selected: number;
}

export interface TableProps {
    title: string
    headers: string[]
    data: string[][]
    pagination: PaginationInfo

    paginationCallback(event: React.MouseEvent<HTMLButtonElement> | null, page : number): Promise<void>
}

const Capitalize = styled.span`text-transform: capitalize`
export const Table: React.FC<TableProps> = ({headers, data, pagination, paginationCallback, title}: TableProps) => {
    return (
        <>
            <TableContainer component={Paper} elevation={2}>
                <Container>
                    <Typography component={"h2"} variant="h6" gutterBottom>
                        <Capitalize>{title}</Capitalize>
                    </Typography>
                </Container>
                <MTable>
                    <TableHead>
                        <TableRow>
                            {headers.map((headerText, k) => {
                                return <TableCell key={k}>{headerText}</TableCell>
                            })}
                        </TableRow>
                    </TableHead>
                    <TableBody>

                        {data.map((row, i) => <TableRow key={i}>
                            {row.map((cell, j) => <TableCell key={j}>
                                {cell}
                            </TableCell>)}
                        </TableRow>)}
                    </TableBody>
                    <TablePagination count={pagination.total}  page={pagination.page} rowsPerPage={pagination.size}  onPageChange={paginationCallback} />
                </MTable>
            </TableContainer>
        </>
    );
};

