import { BsFillTrash2Fill } from "@react-icons/all-files/bs/BsFillTrash2Fill";
import { IconContext } from '@react-icons/all-files';
import { FC } from "react";
import { ToDo } from "../generated";
import { Link } from "react-router-dom";
import { GrDocumentUpdate } from "@react-icons/all-files/gr/GrDocumentUpdate";
import { BiDetail } from "@react-icons/all-files/bi/BiDetail";


interface Props {
  todo: ToDo
}

const TodoItem: FC<Props> = ({ todo }) => {


  return (
    <div className="d-flex text-muted pt-3">

      <div className="pb-3 mb-0 small lh-sm border-bottom w-100">

        <div className="d-flex justify-content-between">
          <Link to={`/detail/${todo.id}`}>
            <IconContext.Provider value={{ size: "20px" }}>
              <BiDetail />
            </IconContext.Provider>
            <strong className="text-gray-dark">{todo.title}</strong>
            <span className="d-block">@{todo.assagin_person}</span>

          </Link>

          <div className="form-check form-switch d-flex align-items-center">
            <input className="form-check-input" type="checkbox" id="flexSwitchCheckDefault" checked={todo.is_complete} />
          </div>

          <div className='d-flex flex-column'>
            <div className='d-flex small  justify-content-end'>
              create:{new Date(todo.created_at * 1000).toLocaleDateString()}
            </div>
            <div className='d-flex small justify-content-end'>
              update:{new Date(todo?.updated_at! * 1000).toLocaleDateString()}
            </div>
          </div>

          <Link to={`/update/${todo.id}`} className='d-flex align-items-center'>
            <IconContext.Provider value={{ size: "20px" }}>
              <GrDocumentUpdate />
            </IconContext.Provider>
          </Link>

          <Link to={`/delete/${todo.id}`} className='d-flex align-items-center'>
            <IconContext.Provider value={{ size: "20px" }}>
              <BsFillTrash2Fill />
            </IconContext.Provider>
          </Link>
        </div>

      </div>
    </div >
  )
}

export default TodoItem
