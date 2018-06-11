// @flow
import * as I from 'immutable'
import * as React from 'react'
import * as Types from '../constants/types/fs'
import {globalStyles} from '../styles'
import {Box, List, Text} from '../common-adapters'
import FolderHeader from './header/container'
import SortBar from './sortbar/container'
import {Still, Editing, Placeholder, Uploading, rowHeight} from './row'
import Footer from './footer/container'

type FolderProps = {
  items: Array<Types.RowItem>,
  path: Types.Path,
  routePath: I.List<string>,
}

class Files extends React.PureComponent<FolderProps> {
  _rowRenderer = (index: number, item: Types.RowItem) => {
    switch (item.rowType) {
      case 'placeholder':
        return <Placeholder key={`placeholder:${item.name}`} />
      case 'still':
        return <Still key={`still:${item.name}`} path={item.path} routePath={this.props.routePath} />
      case 'uploading':
        return <Uploading key={`uploading:${item.name}`} transferID={item.transferID} />
      case 'editing':
        return <Editing key={`editing:${item.name}`} editID={item.editID} routePath={this.props.routePath} />
      default:
        /*::
      let rowType = item.rowType
      declare var ifFlowErrorsHereItsCauseYouDidntHandleAllActionTypesAbove: (rowType: empty) => any
      ifFlowErrorsHereItsCauseYouDidntHandleAllActionTypesAbove(rowType);
      */
        return <Text type="BodyError">This should not happen.</Text>
    }
  }

  render() {
    const content =
      this.props.items && this.props.items.length ? (
        <List fixedHeight={rowHeight} items={this.props.items} renderItem={this._rowRenderer} />
      ) : (
        <Box style={stylesEmptyContainer}>
          <Text type="BodySmall">This is an empty folder.</Text>
        </Box>
      )
    return (
      <Box style={styleOuterContainer}>
        <Box style={stylesContainer}>
          <FolderHeader path={this.props.path} routePath={this.props.routePath} />
          <SortBar path={this.props.path} />
          {content}
          <Footer />
        </Box>
      </Box>
    )
  }
}

const styleOuterContainer = {
  height: '100%',
  position: 'relative',
}

const stylesContainer = {
  ...globalStyles.flexBoxColumn,
  ...globalStyles.fullHeight,
  flex: 1,
}

const stylesEmptyContainer = {
  ...globalStyles.flexBoxColumn,
  ...globalStyles.fullHeight,
  flex: 1,
  justifyContent: 'center',
  alignItems: 'center',
}

export default Files
