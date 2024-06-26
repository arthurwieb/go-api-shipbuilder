
import { GridColDef } from "@mui/x-data-grid";
import "./add.scss"
// import axios from "axios";

type Props = {
    slug: string;
    columns: GridColDef[];
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const Add = (props: Props) => {
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        //FAZER A CHAMADA DA API AQUI @ANDER OU @ARTHUR
        //    axios.post(`http://localhost:3000/${props.slug}`, {        })

    }

    const capsTitle = props.slug[0].toLocaleUpperCase() + props.slug.slice(1);

    return (
        <div className="add">
            <div className="modal">
                <span className="close" onClick={() => props.setOpen(false)}>X</span>
                <h1>Add new {capsTitle}</h1>
                <form onSubmit={handleSubmit}>
                    {props.columns
                        .filter(item => item.field !== "id" && item.field !== "img")
                        .map(column => (
                            <div className="item">
                                <label>{column.headerName}</label>
                                <input type={column.type} placeholder={column.field} />
                            </div>
                        ))}
                    <button>Send</button>
                </form>
            </div>
        </div>
    )
}

export default Add