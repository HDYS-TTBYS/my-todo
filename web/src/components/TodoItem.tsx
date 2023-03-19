import { BsFillTrash2Fill } from "@react-icons/all-files/bs/BsFillTrash2Fill";
import { IconContext } from '@react-icons/all-files';
import { FC } from "react";
import { ToDo } from "../generated";

interface Props {
  todo: ToDo
}

const TodoItem: FC<Props> = ({ todo }) => {


  return (
    <div className="d-flex text-muted pt-3">

      <div className="pb-3 mb-0 small lh-sm border-bottom w-100">
        <div className="d-flex justify-content-between">
          <strong className="text-gray-dark">{todo.title}</strong>

          <div className="form-check form-switch d-flex">
            <input className="form-check-input" type="checkbox" id="flexSwitchCheckDefault" />
          </div>

          <div className='d-flex flex-column'>
            <div className='d-flex small'>
              create:{new Date(todo.created_at * 1000).toLocaleDateString()}
            </div>
            <div className='d-flex small'>
              update:{new Date(todo?.updated_at! * 1000).toLocaleDateString()}
            </div>
          </div>

          <a href="/" className='d-flex'>
            <IconContext.Provider value={{ size: "20px" }}>
              <BsFillTrash2Fill />
            </IconContext.Provider>
          </a>
        </div>
        <span className="d-block">@{todo.assagin_person}</span>
      </div>
    </div>
  )
}

export default TodoItem
