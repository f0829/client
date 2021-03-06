// @flow
import React from 'react'
import {Text} from '../../../../common-adapters'
import UserNotice from '../user-notice'
import {globalColors} from '../../../../styles'

type Props = {
  onShowNewTeamDialog: () => void,
}

const CreateTeamNotice = ({onShowNewTeamDialog}: Props) => (
  <UserNotice username="" bgColor={globalColors.blue4}>
    <Text type="BodySmallSemibold" backgroundMode="Announcements" style={{color: globalColors.black_40}}>
      Make it a team? You'll be able to add and delete members as you wish.
    </Text>
    <Text type="BodySmallPrimaryLink" style={{fontWeight: '600'}} onClick={onShowNewTeamDialog}>
      Enter a team name
    </Text>
  </UserNotice>
)

export default CreateTeamNotice
