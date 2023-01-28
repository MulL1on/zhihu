package draft

type Group struct{}

func (g *Group) Audit() *AuditApi {
	return &insAudit
}
