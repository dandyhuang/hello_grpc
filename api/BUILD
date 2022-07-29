proto_library(
		 name = 'common-proto',
		 srcs = glob([
			 'common/*.proto',
			]),
)

proto_library(
	name = 'abt-proto',
	srcs = glob([
			'abt/*.proto',
		]),
	deps = [
	],
	)
proto_library(
	name = 'zeus-proto',
	srcs = glob([
				'zeus/*.proto',
			]),
	deps = [
			':common-proto',
			':abt-proto',
			'//snapshot_protocol:snapshot_protocol'
		],
)

proto_library(
	name = 'fs-proto',
	srcs = glob([
				'fs/*.proto',
			]),
	deps = [
			'//snapshot_protocol:snapshot_protocol',
			'//protocol/feature_protocol:feature-proto'
		],
	)

proto_library(
	name = 'mixer-proto',
	srcs = glob([
					'mixer/*.proto',
			]),
	deps = [
			':common-proto',
			]
)
proto_library(
	name = 'rank2-proto',
	srcs = glob([
				'rank2/*.proto',
			 ]),
	deps = [
		]
)
proto_library(
	name = 'rank4-proto',
	srcs = glob([
			'rank4/*.proto',
			]),
	deps = [
	]
)

proto_library(
	name = 'rank-proto',
	srcs = glob([
				'rank/*.proto',
	]),
	deps = [
		':common-proto',
		':fs-proto',
		':zeus-proto',
		':mixer-proto',
		':rank4-proto',
	],
)


