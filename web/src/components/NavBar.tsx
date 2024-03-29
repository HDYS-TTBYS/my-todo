import { NavLink } from "react-router-dom";

const NavBar = () => {
  return (
    <div>
      <nav className="navbar navbar-expand navbar-dark bg-dark" aria-label="Second navbar example">
        <div className="container-fluid">
          <a className="navbar-brand" href="/">My Todo</a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarsExample02" aria-controls="navbarsExample02" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>

          <div className="collapse navbar-collapse" id="navbarsExample02">
            <ul className="navbar-nav me-auto">
              <li className="nav-item">
                <NavLink
                  to="/"
                  className={({ isActive, isPending }) =>
                    isPending ? "nav-link" : isActive ? "nav-link active" : "nav-link"
                  }
                >
                  Home
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink
                  to="/create"
                  className={({ isActive, isPending }) =>
                    isPending ? "nav-link" : isActive ? "nav-link active" : "nav-link"
                  }
                >
                  Create
                </NavLink>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>
  )
}

export default NavBar
