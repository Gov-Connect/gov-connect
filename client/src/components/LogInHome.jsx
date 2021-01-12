import React, { Component } from "react";
import axios from "axios";
import RepCard from "./RepCard";
import { CardGroup } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class LogInHome extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user_guid: null,
      reps: [],
    };
  }

  logout = () => {
    localStorage.removeItem("id_token");
    localStorage.removeItem("access_token");
    localStorage.removeItem("profile");
    // location.reload();
  };

  loadReps = () => {
    axios.get(endpoint + "/api/local-reps").then((res) => {
      this.setState({
        user_guid: res.data.user_guid,
        reps: res.data.users_rep_list,
      });
    });
  };

  handleDelete = (guid) => {
    // TODO: We'll need to test this endpoint, but one step at a time
    // axios.delete(endpoint + `/api/localreps/edit?editTask=remove&user_guid=55ee03f2dcd8c8e46b91cbb2e70d9e&rep_guid=${guid}`, {
    //   headers: {
    //     "Content-Type": "application/x-www-form-urlencoded",
    //   },
    // })
    // .then((res) => {
    //     return res;
    // });
    const updatedReps = this.state.reps.filter((rep) => rep.guid !== guid);
    this.setState({
      reps: updatedReps,
    });
  };

  componentDidMount() {
    this.loadReps();
  }

  render() {
    const userList = this.state.reps;
    return (
      <div className="container">
        <br />
        <span className="pull-right">
          <a onClick={this.logout}>Log out</a>
        </span>
        <h2>Open-Gov</h2>
        <p>Hey user</p>
        <div className="row">
          <div className="container">
            <CardGroup>
              {userList &&
                userList.map((localRep, i) => {
                  return (
                    <RepCard
                      key={i}
                      localRep={localRep}
                      deleteRep={this.handleDelete}
                    />
                  );
                })}
            </CardGroup>
          </div>
        </div>
      </div>
    );
  }
}

export default LogInHome;
