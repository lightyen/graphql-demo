import React from "react"
import { Provider } from "react-redux"
import Playground, { store, getSettings, setSettingsString } from "graphql-playground-react"
import { useSelector, useDispatch } from "react-redux"

import "graphql-playground-react/build/static/css/index.css"

const Wrapper = () => {
	const dispatch = useDispatch()
	const settings = useSelector(getSettings)
	React.useEffect(() => {
		if (settings["request.credentials"] !== "same-origin") {
			settings["request.credentials"] = "same-origin"
			dispatch(setSettingsString(JSON.stringify(settings, null, 2)))
		}
	}, [dispatch, settings])
	return <Playground endpoint="/graphql" subscriptionEndpoint="/graphql" />
}

export default () => {
	return (
		<Provider store={store}>
			<Wrapper />
		</Provider>
	)
}
