import React from 'react'
import styles from '../../styles/Home.module.css'
import {PaginationInfo} from "../../types/tasks";
import ReactPaginate from "react-paginate";
import styled from 'styled-components';

const Pagination = styled(ReactPaginate).attrs({
    // You can redifine classes here, if you want.
    activeClassName: 'active', // default to "disabled"
})`
  margin-bottom: 2rem;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  list-style-type: none;
  padding: 0 5rem;

  li a {
    border-radius: 7px;
    padding: 0.1rem 1rem;
    border: gray 1px solid;
    cursor: pointer;
  }

  li.previous a,
  li.next a,
  li.break a {
    border-color: transparent;
  }

  li.active a {
    background-color: #0366d6;
    border-color: transparent;
    color: white;
    min-width: 32px;
  }

  li.disabled a {
    color: grey;
  }

  li.disable,
  li.disabled a {
    cursor: default;
  }
`;

export interface OnPageClickParams {
    selected: number;
}

export interface TableProps {
    title: string
    headers: string[]
    data: string[][]
    pagination: PaginationInfo

    paginationCallback(event: OnPageClickParams): Promise<void>
}

const TableHeading = styled.h1`
  color: var(--platinum);
  text-transform: capitalize;
  width: 100%;
  text-align: center;
`

export const Table: React.FC<TableProps> = ({headers, data, pagination, paginationCallback, title}: TableProps) => {
    const getPagination = () => (<Pagination
        onPageChange={paginationCallback}
        pageCount={pagination.total_pages}
        previousLabel="previous"
        nextLabel="next"
        breakLabel="..."
        breakClassName="page-item"
        breakLinkClassName="page-link"
        pageRangeDisplayed={4}
        marginPagesDisplayed={2}
        containerClassName="pagination justify-content-center"
        pageClassName="page-item"
        pageLinkClassName="page-link"
        previousClassName="page-item"
        previousLinkClassName="page-link"
        nextClassName="page-item"
        nextLinkClassName="page-link"
        activeClassName="active"
        // eslint-disable-next-line no-unused-vars
        // hrefBuilder={(page, pageCount, selected) =>
        //     page >= 1 && page <= pageCount ? `/page/${page}` : '#'
        // }
        // hrefAllControls
    />)
    return (
        <>
            <TableHeading>{title}</TableHeading>
            {getPagination()}
            <div className={styles.table}>
                <div className={`${styles.row} ${styles.header}`}>
                    {headers.map((headerText, k) => {
                        return <div key={k} className={styles.cell}>{headerText}</div>
                    })}
                </div>
                {data.map((row, i) => <div key={i} className={styles.row}>
                    {row.map((cell, j) => <div key={j} className={styles.cell}>
                        {cell}
                    </div>)}
                </div>)}
            </div>
            {getPagination()}
        </>
    );
};

