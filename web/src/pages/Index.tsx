import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';
import { BsFillTrash2Fill } from "@react-icons/all-files/bs/BsFillTrash2Fill";
import { IconContext } from '@react-icons/all-files';


const Index = () => {
  const { data } = useQueryTodos(0)
  const state = {
    curDT: new Date().toLocaleString(),
  }
  return (
    <>
      <NavBar />
      <div className="container">
        <div className="d-flex align-items-center p-3 my-3 text-white bg-purple rounded shadow-sm">
          <div className="lh-1">
            <h1 className="h6 mb-0 text-white lh-1">My Todo</h1>
            <small>Since 2023</small>
          </div>
        </div>

        <div className="my-3 p-3 bg-body rounded shadow-sm">
          <h6 className="border-bottom pb-2 mb-0">Todos</h6>
          <div className="d-flex text-muted pt-3">

            <div className="pb-3 mb-0 small lh-sm border-bottom w-100">
              <div className="d-flex justify-content-between">
                <strong className="text-gray-dark">Full Name</strong>

                <div className="form-check form-switch d-flex">
                  <input className="form-check-input" type="checkbox" id="flexSwitchCheckDefault" />
                </div>

                <div className='d-flex flex-column'>
                  <div>
                    {state.curDT}
                  </div>
                  <div>
                    {state.curDT}
                  </div>
                </div>


                <a href="#">
                  <IconContext.Provider value={{ size: "20px" }}>
                    <BsFillTrash2Fill />
                  </IconContext.Provider>
                </a>
              </div>
              <span className="d-block">@username</span>
            </div>
          </div>
          <small className="d-block text-end mt-3">
            <a href="#">All suggestions</a>
          </small>
        </div>
      </div>

    </>

  )
}

export default Index
