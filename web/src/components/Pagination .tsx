import React, { FC } from 'react'
import { Link, useSearchParams } from 'react-router-dom';

interface Props {
  totalCount: number
}

const Pagination: FC<Props> = ({ totalCount }) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const pageNum = Math.ceil(totalCount / 5)

  return (
    <nav aria-label="Page navigation example">
      {(() => {
        const items = [];
        items.push(<li key={0} className="page-item">
          <a href={"/"} className="page-link" aria-label="Previous">
            <span aria-hidden="true">&laquo;</span>
          </a>
        </li>)
        for (let i = 0; i < pageNum; i++) {
          items.push(<li key={i + 1} className={Number(searchParams.get("offset")) / 5 === i ? "page-item active" : "page-item"}>
            <a href={`/?offset=${i * 5}`} className="page-link">{i + 1}</a>
          </li>)
        }
        items.push(<li key={pageNum + 1} className="page-item">
          <a href={`/?offset=${pageNum * 5 - 5}`} className="page-link" aria-label="Next">
            <span aria-hidden="true">&raquo;</span>
          </a>
        </li>)
        return <ul className="pagination justify-content-end">{items}</ul>;
      })()}
    </nav >
  )
}

export default Pagination 
