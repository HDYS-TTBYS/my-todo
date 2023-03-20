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
    <tr>
      <td>
        <Link to={`/detail/${todo.id}`}>
          <div className="d-flex">
            <IconContext.Provider value={{ size: "20px" }}>
              <BiDetail />
            </IconContext.Provider>
            <span className="d-block small">@{todo.assagin_person}</span>
          </div>

          <div className="container">
            <strong className="text-gray-dark d-flex small">{todo.title}</strong>
          </div>
        </Link>
      </td>

      <td>
        <div className="form-check form-switch d-flex align-items-center justify-content-center">
          <input className="form-check-input" type="checkbox" disabled id="flexSwitchCheckDefault" checked={todo.is_complete} />
        </div>
      </td>

      <td>
        <Link to={`/update/${todo.id}`} className='d-flex align-items-center justify-content-center'>
          <IconContext.Provider value={{ size: "20px" }}>
            <GrDocumentUpdate />
          </IconContext.Provider>
        </Link>
      </td>

      <td>
        <div className='d-flex small justify-content-end'>
          {new Date(todo.created_at * 1000).toLocaleDateString()}
        </div>
      </td>

      <td>
        <Link to={`/delete/${todo.id}`} className='d-flex align-items-center justify-content-center'>
          <IconContext.Provider value={{ size: "20px" }}>
            <BsFillTrash2Fill />
          </IconContext.Provider>
        </Link>
      </td>
    </tr >
  )
}

export default TodoItem
