import NavBar from '../components/NavBar';
import { useQueryTodos } from '../hooks/useQueryTotos';
import { BsFillTrash2Fill } from "@react-icons/all-files/bs/BsFillTrash2Fill";
import { IconContext } from '@react-icons/all-files';
import { time } from 'console';


const Index = () => {
  const { data } = useQueryTodos(0)
  const state = {
    curDT: new Date().toLocaleString(),
  }
  return (
    <>
      <NavBar />
      <div className='container'>
        {JSON.stringify(data)}
        <div className="my-3 p-3 bg-body rounded shadow-sm">
          <h6 className="border-bottom pb-2 mb-0 fs-2">Todos</h6>

          <div className="d-flex align-items-center text-muted pt-3 justify-content-between border-bottom">
            <div className="pb-3 mb-0 small lh-sm w-100">
              <div className="d-flex">
                <strong className="text-gray-dark fs-4">Title</strong>
              </div>
              <span className="d-block fs-5">@assaginPerson:</span>
            </div>

            <div className="form-check form-switch d-flex align-items-center w-100 flex-column justify-content-center">
              <input className="form-check-input d-flex" type="checkbox" id="flexSwitchCheckChecked" checked />
              <label className="form-check-label d-flex" htmlFor="flexSwitchCheckChecked">completed</label>
            </div>

            <div className="pb-3 mb-0 small lh-sm w-100 d-flex flex-column justify-content-center">
              <div className="d-flex">
                created: {state.curDT}
              </div>
              <div className="d-flex">
                updated: {state.curDT}
              </div>
            </div>

            <div className='pb-3 mb-0 small lh-sm w-100 d-flex justify-content-center'>
              <IconContext.Provider value={{ size: "50px" }}>
                <BsFillTrash2Fill />
              </IconContext.Provider>
            </div>

          </div>

        </div>
      </div>
    </>

  )
}

export default Index
