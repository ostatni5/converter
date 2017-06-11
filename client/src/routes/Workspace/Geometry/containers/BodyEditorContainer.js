/* @flow */

import React from 'react';
import { connect } from 'react-redux';
import { t } from 'i18n';
import type { Body, BodyGeometry, ConstructionPath } from '../../model';
import type { GeometryType } from '../../bodyModel';
import BodyEditorLayout from '../components/BodyEditorLayout';
import { defaultBodyForType } from '../defaults';
import { actionCreator } from '../../reducer';

type Props = {
  bodyId: number,
  bodyGeometry: BodyGeometry,
  setGeometry: (geometry: BodyGeometry) => void,
  constructionPath: ConstructionPath & { action: 'update' | 'create' },
  closeModal: () => void,

  updateBody: (body: Body) => void,
  createBodyInZone: (body: Body, path: ConstructionPath) => void,
}

class BodyEditorContainer extends React.Component {
  props: Props

  typeUpdate = (type: GeometryType) => {
    const newGeometry = { type, ...defaultBodyForType(type) };
    this.props.setGeometry(newGeometry);
  }

  geometryUpdate = (field: string, value: Object) => {
    const newGeometry: Object = { ...this.props.bodyGeometry, [field]: value };
    this.props.setGeometry((newGeometry: BodyGeometry));
  }

  applyChanges = () => {
    if (this.props.bodyId !== undefined) {
      const body: Body = { id: this.props.bodyId, geometry: this.props.bodyGeometry };
      this.props.updateBody(body);
    } else {
      const body: Body = { id: this.props.bodyId, geometry: this.props.bodyGeometry };
      this.props.createBodyInZone(body, this.props.constructionPath);
    }
    this.props.closeModal();
  }


  render() {
    return (
      <BodyEditorLayout
        bodyGeometry={this.props.bodyGeometry}
        typeUpdate={this.typeUpdate}
        geometryUpdate={this.geometryUpdate}
        submit={this.applyChanges}
        submitBtnName={
          this.props.bodyId !== undefined
            ? t('workspace.editor.updateBtn')
            : t('workspace.editor.createBtn')
        }
      />
    );
  }
}

const mapDispatchToProps = (dispatch) => {
  return {
    createBodyInZone: (body, path) => dispatch(actionCreator.createBodyInZone(body, path)),
    updateBody: body => dispatch(actionCreator.updateBody(body)),
  };
};

export default connect(
  undefined,
  mapDispatchToProps,
)(BodyEditorContainer);
