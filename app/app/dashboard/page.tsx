import VoteList from "#/ui/vote/vote-list"
  
export default async function DashboardPage() {
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <VoteList />
    </div>
  );
}