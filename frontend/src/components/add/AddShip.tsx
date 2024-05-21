
import { GridColDef } from "@mui/x-data-grid";
import "./add.scss"
import { useForm, SubmitHandler } from "react-hook-form"
//import { URL } from "../../api/api";
//import axios from "axios";

type Props = {
    slug: string;
    columns: GridColDef[];
    setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

type Ship = {
    id: number;
    shipName: string;
    shipType: number;
}

const Add = (props: Props) => {

    const { register, handleSubmit } = useForm<Ship>();
    const onSubmit: SubmitHandler<Ship> = (data) => {
        console.log(data)
        props.setOpen(false)


        //axios.post(`${URL}ship`, {data})
        //history.forward("/user/ship/id")
    }


    // const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    //     e.preventDefault();

    //     //FAZER A CHAMADA DA API AQUI @ANDER OU @ARTHUR
    //     //    axios.post(`http://localhost:3000/${props.slug}`, {        })

    // }

    return (
        <div className="add">
            <div className="modal">
                <span className="close" onClick={() => props.setOpen(false)}>X</span>
                <h1>Add new Ship</h1>
                    {<form onSubmit={handleSubmit(onSubmit)}>
                        <div className="item">
                            <label>Ship Name</label>
                            <input type="shipName" placeholder="Ship Name" {...register("shipName")}/>
                        </div>
                        <div className="item">
                            <label>Ship Type</label>
                             <input type="shipType" placeholder="Ship Type" {...register("shipType")}/>
                        </div>
                        <button>Send</button>
                    </form>
                
                
                /* <form onSubmit={handleSubmit}>
                    {props.columns
                        .filter(item => item.field !== "id" && item.field !== "img")
                        .map(column => (
                            <div className="item">
                                <label>{column.headerName}</label>
                                <input type={column.type} placeholder={column.field} />
                            </div>
                        ))}
                    <button>Send</button>
                </form> */}
            </div>
        </div>
    )
}

export default Add