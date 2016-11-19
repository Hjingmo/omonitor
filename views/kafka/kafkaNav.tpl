{{define "headbtn"}}
					<a href="/kafka/consumers?env={{.env}}" class="btn btn-sm btn-primary "> Consumer Groups </a>
					<a href="/kafka/topics?env={{.env}}" class="btn btn-sm btn-primary "> Topic List </a>
					<a href="/kafka/servers?env={{.env}}" class="btn btn-sm btn-primary "> Kafka Cluster </a>
{{end}}