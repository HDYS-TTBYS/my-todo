import NavBar from "../components/NavBar";

export default function ErrorPage() {


  return (
    <>
      <NavBar />
      <div className="container">
        <div id="d-flex align-items-center justify-content-center">
          <h1 className="d-flex mt-5 align-items-center justify-content-center">申し訳ございません。</h1>
          <p className="d-flex mt-3 align-items-center justify-content-center">お探しのページは見つかりませんでした。</p>
        </div>
      </div>
    </>
  );
}
