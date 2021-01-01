import React, { Component } from "react";
import { Card, Icon } from "semantic-ui-react";

class RepCard extends Component {
  deleteRep = (id) => {
    this.props.deleteRep(id);
  };

  render() {
    const { localRep } = this.props;

    const footer = (
      <a>
        <Icon
          name="delete"
          color="red"
          onClick={() => this.deleteRep(localRep.guid)}
        />
        Delete Rep
      </a>
    );

    const description = (
      <div className="panel">
        <p>Office: {localRep.office}</p>
        <p>Percent of Votes Missed: {localRep.percent_missed_votes}%</p>
        <p>
          <a href={localRep.gov_web}>Goverment Web Page</a>
        </p>
        <p>
          <a href={`https://www.twitter.com/${localRep.twitter}`}>Twitter</a>
        </p>
      </div>
    );

    return (
      <Card
        image="https://react.semantic-ui.com/images/avatar/large/elliot.jpg"
        header={localRep.name + " " + localRep.LastName}
        meta={localRep.location}
        description={description}
        extra={footer}
      />
    );
  }
}

export default RepCard;
